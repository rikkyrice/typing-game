package infrastructure

import (
	"api/db"
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
)

const selectWordByIDQuery string = `
	SELECT * FROM words WHERE id = $1
`

const selectWordByWordListIDQuery string = `
	SELECT * FROM words WHERE word_list_id = $1
`

const insertWordQuery string = `
	INSERT INTO words
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
`

const updateWordByIDQuery string = `
	UPDATE words
	SET
		word = $1,
		meaning = $2,
		explanation = $3,
		updated_at = $4
	WHERE id = $5
`

const deleteWordByIDQuery string = `
	DELETE FROM words WHERE id = $1
`

const deleteAllWordByWordListIDQuery string = `
	DELETE FROM words WHERE word_list_id = $1
`

// NewWordRepository 単語帳リポジトリの生成
func NewWordRepository(conn *db.DBConn) (repository.WordRepository, error) {
	errs := []error{}

	selectWordByIDPstmt, err := conn.GetPstmt(selectWordByIDQuery)
	errs = append(errs, err)

	selectWordByWordListIDPstmt, err := conn.GetPstmt(selectWordByWordListIDQuery)
	errs = append(errs, err)

	insertWordPstmt, err := conn.GetPstmt(insertWordQuery)
	errs = append(errs, err)

	updateWordByIDPstmt, err := conn.GetPstmt(updateWordByIDQuery)
	errs = append(errs, err)

	deleteWordByIDPstmt, err := conn.GetPstmt(deleteWordByIDQuery)
	errs = append(errs, err)

	deleteAllWordByWordListIDPstmt, err := conn.GetPstmt(deleteAllWordByWordListIDQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrapf(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &wordRepository{
		selectWordByIDPstmt:            selectWordByIDPstmt,
		selectWordByWordListIDPstmt:    selectWordByWordListIDPstmt,
		insertWordPstmt:                insertWordPstmt,
		updateWordByIDPstmt:            updateWordByIDPstmt,
		deleteWordByIDPstmt:            deleteWordByIDPstmt,
		deleteAllWordByWordListIDPstmt: deleteAllWordByWordListIDPstmt,
	}, nil
}

// wordRepository 単語帳リポジトリインターフェース
type wordRepository struct {
	selectWordByIDPstmt            *sql.Stmt
	selectWordByWordListIDPstmt    *sql.Stmt
	insertWordPstmt                *sql.Stmt
	updateWordByIDPstmt            *sql.Stmt
	deleteWordByIDPstmt            *sql.Stmt
	deleteAllWordByWordListIDPstmt *sql.Stmt
}

func (wR *wordRepository) FindWordByID(id string) (*model.Word, *apierror.Error) {
	var w model.Word

	if err := wR.selectWordByIDPstmt.QueryRow(id).Scan(&w.ID, &w.WordListID, &w.Word, &w.Yomi, &w.Meaning, &w.MYomi, &w.Explanation, &w.IsRemembered, &w.CreatedAt, &w.UpdatedAt); err != nil {
		return nil, apierror.NewError(http.StatusNotFound, errors.Wrapf(err, "ID[%s]の単語が見つかりません。", id))
	}
	return &w, nil
}

func (wR *wordRepository) FindWordByWordListID(wlID string) ([]*model.Word, *apierror.Error) {
	ws := []*model.Word{}

	rows, err := wR.selectWordByWordListIDPstmt.Query(wlID)
	if err != nil {
		return ws, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "クエリの実行に失敗しました。"))
	}

	for rows.Next() {
		var w model.Word
		if err := rows.Scan(&w.ID, &w.WordListID, &w.Word, &w.Yomi, &w.Meaning, &w.MYomi, &w.Explanation, &w.IsRemembered, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "レコードの読み取りに失敗しました。"))
		}
		ws = append(ws, &w)
	}

	return ws, nil
}

func (wR *wordRepository) CreateWord(w model.Word) (*model.Word, *apierror.Error) {
	_, err := wR.insertWordPstmt.Exec(&w.ID, &w.WordListID, &w.Word, &w.Yomi, &w.Meaning, &w.MYomi, &w.Explanation, &w.IsRemembered, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "単語の作成に失敗しました。"))
	}
	return &model.Word{
		ID:          w.ID,
		WordListID:  w.WordListID,
		Word:        w.Word,
		Meaning:     w.Meaning,
		Explanation: w.Explanation,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}, nil
}

func (wR *wordRepository) CreateAllWord(ws []model.Word) ([]*model.Word, *apierror.Error) {
	words := []*model.Word{}
	for _, w := range ws {
		word, err := wR.CreateWord(w)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

func (wR *wordRepository) UpdateWordByID(id string, w model.Word) (*model.Word, *apierror.Error) {
	res, err := wR.updateWordByIDPstmt.Exec(w.Word, w.Meaning, w.Explanation, w.UpdatedAt, id)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "単語の更新に失敗しました。"))
	}
	rec, err := res.RowsAffected()
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "更新した列の件数の取得に失敗しました。"))
	}
	if rec == 0 {
		return nil, apierror.NewError(http.StatusNotFound, errors.New("更新対象のリソースが存在しません。"))
	}
	return &model.Word{
		ID:           id,
		WordListID:   w.WordListID,
		Word:         w.Word,
		Yomi:         w.Yomi,
		Meaning:      w.Meaning,
		MYomi:        w.MYomi,
		Explanation:  w.Explanation,
		IsRemembered: w.IsRemembered,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}, nil
}

func (wR *wordRepository) RemoveWordByID(id string) *apierror.Error {
	res, err := wR.deleteWordByIDPstmt.Exec(id)
	if err != nil {
		return apierror.NewError(http.StatusNotFound, errors.Wrapf(err, "ID[%s]の単語の削除に失敗しました。", id))
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

func (wR *wordRepository) RemoveAllWordByWordListID(wlID string) *apierror.Error {
	ws, err := wR.FindWordByWordListID(wlID)
	if err != nil {
		return err
	}
	for _, w := range ws {
		err := wR.RemoveWordByID(w.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
