package infrastructure

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"
	"net/http"
	"sort"

	"api/db"

	"github.com/pkg/errors"
)

const selectWordListByIDQuery string = `
	SELECT * FROM wordlistsSummaries WHERE id = $1
`

const selectWordListByUserIDQuery string = `
	SELECT * FROM wordlistsSummaries WHERE user_id = $1
`

const insertWordListQuery string = `
	INSERT INTO wordlists
	VALUES($1,$2,$3,$4,$5,$6)
`

const updateWordListByIDQuery string = `
	UPDATE wordlists
	SET
		word_list_title = $1,
		explanation = $2,
		updated_at = $3
	WHERE id = $4
`

const deleteWordListByIDQuery string = `
	DELETE FROM wordlists WHERE id = $1
`

// NewWordListRepository 単語帳リポジトリの生成
func NewWordListRepository(conn *db.DBConn) (repository.WordListRepository, error) {
	errs := []error{}

	selectWordListByIDPstmt, err := conn.GetPstmt(selectWordListByIDQuery)
	errs = append(errs, err)

	selectWordListByUserIDPstmt, err := conn.GetPstmt(selectWordListByUserIDQuery)
	errs = append(errs, err)

	insertWordListPstmt, err := conn.GetPstmt(insertWordListQuery)
	errs = append(errs, err)

	updateWordListByIDPstmt, err := conn.GetPstmt(updateWordListByIDQuery)
	errs = append(errs, err)

	deleteWordListByIDPstmt, err := conn.GetPstmt(deleteWordListByIDQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrap(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &wordListRepository{
		selectWordListByIDPstmt:     selectWordListByIDPstmt,
		selectWordListByUserIDPstmt: selectWordListByUserIDPstmt,
		insertWordListPstmt:         insertWordListPstmt,
		updateWordListByIDPstmt:     updateWordListByIDPstmt,
		deleteWordListByIDPstmt:     deleteWordListByIDPstmt,
	}, nil
}

// wordListRepository 単語帳リポジトリインターフェース
type wordListRepository struct {
	selectWordListByIDPstmt     *sql.Stmt
	selectWordListByUserIDPstmt *sql.Stmt
	insertWordListPstmt         *sql.Stmt
	updateWordListByIDPstmt     *sql.Stmt
	deleteWordListByIDPstmt     *sql.Stmt
}

func (wlR *wordListRepository) FindWordListByID(id string) (*model.WordListSummary, *apierror.Error) {
	var wl model.WordListSummary

	if err := wlR.selectWordListByIDPstmt.QueryRow(id).Scan(&wl.ID, &wl.UserID, &wl.Title, &wl.Explanation, &wl.WordCount, &wl.PlayCount, &wl.PlayedAt, &wl.CreatedAt, &wl.UpdatedAt); err != nil {
		return nil, apierror.NewError(http.StatusNotFound, errors.Wrap(err, "単語帳が見つかりません。"))
	}
	return &wl, nil
}

func (wlR *wordListRepository) FindWordListByUserID(userID string) ([]*model.WordListSummary, *apierror.Error) {
	wls := []*model.WordListSummary{}

	rows, err := wlR.selectWordListByUserIDPstmt.Query(userID)
	if err != nil {
		return wls, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "クエリの実行に失敗しました。"))
	}

	for rows.Next() {
		var wl model.WordListSummary
		if err := rows.Scan(&wl.ID, &wl.UserID, &wl.Title, &wl.Explanation, &wl.WordCount, &wl.PlayCount, &wl.PlayedAt, &wl.CreatedAt, &wl.UpdatedAt); err != nil {
			return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "レコードの読み取りに失敗しました。"))
		}
		wls = append(wls, &wl)
	}

	sort.Slice(wls, func(i, j int) bool { return wls[i].CreatedAt.Before(wls[j].CreatedAt) })

	return wls, nil
}

func (wlR *wordListRepository) CreateWordList(wl model.WordList) (*model.WordList, *apierror.Error) {
	_, err := wlR.insertWordListPstmt.Exec(wl.ID, wl.UserID, wl.Title, wl.Explanation, wl.CreatedAt, wl.UpdatedAt)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "単語帳の作成に失敗しました。"))
	}
	return &model.WordList{
		ID:          wl.ID,
		UserID:      wl.UserID,
		Title:       wl.Title,
		Explanation: wl.Explanation,
		CreatedAt:   wl.CreatedAt,
		UpdatedAt:   wl.UpdatedAt,
	}, nil
}

func (wlR *wordListRepository) UpdateWordListByID(id string, wl model.WordList) (*model.WordList, *apierror.Error) {
	res, err := wlR.updateWordListByIDPstmt.Exec(wl.Title, wl.Explanation, wl.UpdatedAt, id)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "単語帳の更新に失敗しました。"))
	}
	rec, err := res.RowsAffected()
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "更新した列の件数の取得に失敗しました。"))
	}
	if rec == 0 {
		return nil, apierror.NewError(http.StatusNotFound, errors.New("更新対象のリソースが存在しません。"))
	}
	return &model.WordList{
		ID:          id,
		UserID:      wl.UserID,
		Title:       wl.Title,
		Explanation: wl.Explanation,
		CreatedAt:   wl.CreatedAt,
		UpdatedAt:   wl.UpdatedAt,
	}, nil
}

func (wlR *wordListRepository) RemoveWordListByID(id string) *apierror.Error {
	res, err := wlR.deleteWordListByIDPstmt.Exec(id)
	if err != nil {
		return apierror.NewError(http.StatusNotFound, errors.Wrap(err, "単語帳の削除に失敗しました。"))
	}
	rec, err := res.RowsAffected()
	if err != nil {
		return apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "削除した列の件数の取得に失敗しました。"))
	}
	if rec == 0 {
		return apierror.NewError(http.StatusNotFound, errors.New("削除対象のリソースが存在しません。"))
	}
	return nil
}
