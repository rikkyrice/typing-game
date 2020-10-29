package infrastructure

import (
	"api/db"
	"api/internal/common/util"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"

	"github.com/pkg/errors"
)

const selectScoreByWordListIDQuery string = `
	SELECT * FROM scores WHERE word_list_id = ?
`

const selectLatestScoreByWordListIDQuery string = `
	SELECT * 
	FROM scores 
	WHERE word_list_id = ?
	ORDER BY played_at desc
	FETCH FIRST 1 ROWS ONLY
`

const insertScoreQuery string = `
	INSERT INTO scores
	VALUES(?,?,?,?,?,?,?)
`

const deleteScoreByIDQuery string = `
	DELETE FROM scores WHERE id = ?
`

const deleteAllScoreByWordListIDQuery string = `
	DELETE FROM scores WHERE word_list_id = ?
`

// NewScoreRepository スコアリポジトリの生成
func NewScoreRepository(conn *db.DBConn) (repository.ScoreRepository, error) {
	errs := []error{}

	selectScoreByWordListIDPstmt, err := conn.GetPstmt(selectScoreByWordListIDQuery)
	errs = append(errs, err)

	selectLatestScoreByWordListIDPstmt, err := conn.GetPstmt(selectLatestScoreByWordListIDQuery)
	errs = append(errs, err)

	insertScorePstmt, err := conn.GetPstmt(insertScoreQuery)
	errs = append(errs, err)

	deleteScoreByIDPstmt, err := conn.GetPstmt(deleteScoreByIDQuery)
	errs = append(errs, err)

	deleteAllScoreByWordListIDPstmt, err := conn.GetPstmt(deleteAllScoreByWordListIDQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrap(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &scoreRepository{
		selectScoreByWordListIDPstmt:       selectScoreByWordListIDPstmt,
		selectLatestScoreByWordListIDPstmt: selectLatestScoreByWordListIDPstmt,
		insertScorePstmt:                   insertScorePstmt,
		deleteScoreByIDPstmt:               deleteScoreByIDPstmt,
		deleteAllScoreByWordListIDPstmt:    deleteAllScoreByWordListIDPstmt,
	}, nil
}

// scoreRepository スコアリポジトリインターフェース
type scoreRepository struct {
	selectScoreByWordListIDPstmt       *sql.Stmt
	selectLatestScoreByWordListIDPstmt *sql.Stmt
	insertScorePstmt                   *sql.Stmt
	deleteScoreByIDPstmt               *sql.Stmt
	deleteAllScoreByWordListIDPstmt    *sql.Stmt
}

func (sR *scoreRepository) FindScoreByWordListID(wlID string) ([]*model.Score, error) {
	ss := []*model.Score{}

	rows, err := sR.selectScoreByWordListIDPstmt.Query(wlID)
	if err != nil {
		return ss, errors.Wrap(err, "クエリ実行に失敗")
	}

	for rows.Next() {
		var s model.Score
		if err := rows.Scan(&s.ID, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.CorrectRate, &s.PlayedAt); err != nil {
			return nil, errors.Wrap(err, "クエリの読み込みに失敗")
		}
		ss = append(ss, &s)
	}

	return ss, nil
}

func (sR *scoreRepository) FIndLatestScoreByWordListID(wlID string) (*model.Score, error) {
	var s model.Score

	if err := sR.selectLatestScoreByWordListIDPstmt.QueryRow(wlID).Scan(&s.ID, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.CorrectRate, &s.PlayedAt); err != nil {
		return nil, errors.Wrap(err, "クエリ実行に失敗")
	}
	return &s, nil
}

func (sR *scoreRepository) CreateScore(s model.Score) (*model.Score, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	_, err = sR.insertScorePstmt.Exec(id, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.CorrectRate, &s.PlayedAt)
	if err != nil {
		return nil, errors.Wrap(err, "スコアの作成に失敗しました。")
	}
	return &model.Score{
		ID:             id,
		WordListID:     s.WordListID,
		PlayCount:      s.PlayCount,
		ClearTypeCount: s.ClearTypeCount,
		MissTypeCount:  s.MissTypeCount,
		CorrectRate:    s.CorrectRate,
		PlayedAt:       s.PlayedAt,
	}, err
}

func (sR *scoreRepository) RemoveScoreByID(id string) error {
	_, err := sR.deleteScoreByIDPstmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "スコアの削除に失敗しました。")
	}
	return nil
}

func (sR *scoreRepository) RemoveAllScoreByWordListID(wlID string) error {
	_, err := sR.deleteAllScoreByWordListIDPstmt.Exec(wlID)
	if err != nil {
		return errors.Wrap(err, "スコアの全削除に失敗しました。")
	}
	return nil
}
