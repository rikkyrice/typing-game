package infrastructure

import (
	"api/internal/common/util"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"

	"api/db"

	"github.com/pkg/errors"
)

const selectWordListByIDQuery string = `
	SELECT * FROM wordlists WHERE id = ?
`

const selectWordListByUserIDQuery string = `
	SELECT * FROM wordlists WHERE user_id = ?
`

const insertWordListQuery string = `
	INSERT INTO wordlists
	VALUES(?,?,?,?,?,?)
`

const updateWordListByIDQuery string = `
	UPDATE wordlists
	SET
		word_list_title = ?,
		explanation = ?,
		updated_at = ?
	WHERE id = ?
`

const deleteWordListByIDQuery string = `
	DELETE FROM wordlists WHERE id = ?
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
			return nil, errors.Wrapf(err, "ステートメントの作成に失敗しました。")
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

func (wlR *wordListRepository) FindWordListByID(id string) (*model.WordList, error) {
	var wl model.WordList

	if err := wlR.selectWordListByIDPstmt.QueryRow(id).Scan(&wl.ID, &wl.UserID, &wl.Title, &wl.Explanation, &wl.CreatedAt, &wl.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "クエリ実行に失敗")
	}
	return &wl, nil
}

func (wlR *wordListRepository) FindWordListByUserID(userID string) ([]*model.WordList, error) {
	wls := []*model.WordList{}

	rows, err := wlR.selectWordListByUserIDPstmt.Query(userID)
	if err != nil {
		return wls, errors.Wrap(err, "クエリ実行に失敗")
	}

	for rows.Next() {
		var wl model.WordList
		if err := rows.Scan(&wl.ID, &wl.UserID, &wl.Title, &wl.Explanation, &wl.CreatedAt, &wl.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "クエリの読み込みに失敗")
		}
		wls = append(wls, &wl)
	}

	return wls, nil
}

func (wlR *wordListRepository) CreateWordList(wl model.WordList) (*model.WordList, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	_, err = wlR.insertWordListPstmt.Exec(id, wl.UserID, wl.Title, wl.Explanation, wl.CreatedAt, wl.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "単語帳の作成に失敗しました。")
	}
	return &model.WordList{
		ID:          id,
		UserID:      wl.UserID,
		Title:       wl.Title,
		Explanation: wl.Explanation,
		CreatedAt:   wl.CreatedAt,
		UpdatedAt:   wl.UpdatedAt,
	}, err
}

func (wlR *wordListRepository) UpdateWordListByID(id string, wl model.WordList) (*model.WordList, error) {
	_, err := wlR.updateWordListByIDPstmt.Exec(wl.Title, wl.Explanation, wl.UpdatedAt, id)
	if err != nil {
		return nil, errors.Wrap(err, "単語帳の更新に失敗しました。")
	}
	return &model.WordList{
		ID:          id,
		UserID:      wl.UserID,
		Title:       wl.Title,
		Explanation: wl.Explanation,
		CreatedAt:   wl.CreatedAt,
		UpdatedAt:   wl.UpdatedAt,
	}, err
}

func (wlR *wordListRepository) RemoveWordListByID(id string) error {
	_, err := wlR.deleteWordListByIDPstmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "単語帳の削除に失敗しました。")
	}
	return nil
}
