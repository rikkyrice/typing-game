package infrastructure

import (
	"api/db"
	"api/internal/common/util"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"

	"github.com/pkg/errors"
)

const selectWordByIDQuery string = `
	SELECT * FROM words WHERE id = ?
`

const selectWordByWordListIDQuery string = `
	SELECT * FROM words WHERE word_list_id = ?
`

const insertWordQuery string = `
	INSERT INTO words
	VALUES(?,?,?,?,?,?,?)
`

const updateWordByIDQuery string = `
	UPDATE words
	SET
		word__title = ?,
		explanation = ?,
		updated_at = ?
	WHERE id = ?
`

const deleteWordByIDQuery string = `
	DELETE FROM words WHERE id = ?
`

const deleteAllWordByWordListIDQuery string = `
	DELETE FROM words WHERE word_list_id = ?
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

func (wR *wordRepository) FindWordByID(id string) (*model.Word, error) {
	var w model.Word

	if err := wR.selectWordByIDPstmt.QueryRow(id).Scan(&w.ID, &w.WordListID, &w.Word, &w.Meaning, &w.Explanation, &w.CreatedAt, &w.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "クエリ実行に失敗")
	}
	return &w, nil
}

func (wR *wordRepository) FindWordByWordListID(wlID string) ([]*model.Word, error) {
	ws := []*model.Word{}

	rows, err := wR.selectWordByWordListIDPstmt.Query(wlID)
	if err != nil {
		return ws, errors.Wrap(err, "クエリ実行に失敗")
	}

	for rows.Next() {
		var w model.Word
		if err := rows.Scan(&w.ID, &w.WordListID, &w.Word, &w.Meaning, &w.Explanation, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "クエリの読み込みに失敗")
		}
		ws = append(ws, &w)
	}

	return ws, nil
}

func (wR *wordRepository) CreateWord(w model.Word) (*model.Word, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	_, err = wR.insertWordPstmt.Exec(id, &w.WordListID, &w.Word, &w.Meaning, &w.Explanation, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "単語の作成に失敗しました。")
	}
	return &model.Word{
		ID:          id,
		WordListID:  w.WordListID,
		Word:        w.Word,
		Meaning:     w.Meaning,
		Explanation: w.Explanation,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}, err
}

func (wR *wordRepository) CreateAllWord(ws []model.Word) ([]*model.Word, error) {
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

func (wR *wordRepository) UpdateWordByID(id string, w model.Word) (*model.Word, error) {
	_, err := wR.updateWordByIDPstmt.Exec(w.Word, w.Meaning, w.Explanation, w.UpdatedAt, id)
	if err != nil {
		return nil, errors.Wrap(err, "単語の更新に失敗しました。")
	}
	return &model.Word{
		ID:          id,
		WordListID:  w.WordListID,
		Word:        w.Word,
		Meaning:     w.Meaning,
		Explanation: w.Explanation,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}, err
}

func (wR *wordRepository) RemoveWordByID(id string) error {
	_, err := wR.deleteWordByIDPstmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "単語の削除に失敗しました。")
	}
	return nil
}

func (wR *wordRepository) RemoveAllWordByWordListID(wlID string) error {
	_, err := wR.deleteAllWordByWordListIDPstmt.Exec(wlID)
	if err != nil {
		return errors.Wrap(err, "単語の全削除に失敗しました。")
	}
	return nil
}
