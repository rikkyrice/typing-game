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

const selectScoreByWordListIDQuery string = `
	SELECT * FROM scores WHERE word_list_id = $1
`

const selectLatestScoreByWordListIDQuery string = `
	SELECT * 
	FROM scores 
	WHERE word_list_id = $1
	ORDER BY played_at desc
	FETCH FIRST 1 ROWS ONLY
`

const insertScoreQuery string = `
	INSERT INTO scores
	VALUES($1,$2,$3,$4,$5,$6)
`

const deleteScoreByIDQuery string = `
	DELETE FROM scores WHERE id = $1
`

const deleteAllScoreByWordListIDQuery string = `
	DELETE FROM scores WHERE word_list_id = $1
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

func (sR *scoreRepository) FindScoreByWordListID(wlID string) ([]*model.Score, *apierror.Error) {
	ss := []*model.Score{}

	rows, err := sR.selectScoreByWordListIDPstmt.Query(wlID)
	if err != nil {
		return ss, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "クエリの実行に失敗しました。"))
	}

	for rows.Next() {
		var s model.Score
		if err := rows.Scan(&s.ID, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.PlayedAt); err != nil {
			return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "レコードの読み取りに失敗しました。"))
		}
		ss = append(ss, &s)
	}

	return ss, nil
}

func (sR *scoreRepository) FIndLatestScoreByWordListID(wlID string) (*model.Score, *apierror.Error) {
	var s model.Score

	if err := sR.selectLatestScoreByWordListIDPstmt.QueryRow(wlID).Scan(&s.ID, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.PlayedAt); err != nil {
		return nil, apierror.NewError(http.StatusNotFound, errors.Wrapf(err, "ID[%s]の単語帳の最新のスコアが見つかりません。", wlID))
	}
	return &s, nil
}

func (sR *scoreRepository) CreateScore(s model.Score) (*model.Score, *apierror.Error) {
	_, err := sR.insertScorePstmt.Exec(&s.ID, &s.WordListID, &s.PlayCount, &s.ClearTypeCount, &s.MissTypeCount, &s.PlayedAt)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "スコアの作成に失敗しました。"))
	}
	return &model.Score{
		ID:             s.ID,
		WordListID:     s.WordListID,
		PlayCount:      s.PlayCount,
		ClearTypeCount: s.ClearTypeCount,
		MissTypeCount:  s.MissTypeCount,
		PlayedAt:       s.PlayedAt,
	}, nil
}

func (sR *scoreRepository) RemoveScoreByID(id string) *apierror.Error {
	res, err := sR.deleteScoreByIDPstmt.Exec(id)
	if err != nil {
		return apierror.NewError(http.StatusNotFound, errors.Wrapf(err, "ID[%s]のスコアの削除に失敗しました。", id))
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

func (sR *scoreRepository) RemoveAllScoreByWordListID(wlID string) *apierror.Error {
	ss, err := sR.FindScoreByWordListID(wlID)
	if err != nil {
		return err
	}
	for _, s := range ss {
		err := sR.RemoveScoreByID(s.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
