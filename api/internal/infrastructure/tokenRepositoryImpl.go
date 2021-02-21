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

const selectTokenByUserIDQuery string = `
	SELECT *
	FROM token
	WHERE user_id = $1
	ORDER BY created_at desc
	FETCH FIRST 1 ROW ONLY
`

const insertTokenQuery string = `
	INSERT INTO token
	VALUES($1,$2,$3,$4)
`

const deleteTokenByUserIDQuery string = `
	DELETE FROM token
	WHERE user_id = $1
`

// NewTokenRepository ORMapper
func NewTokenRepository(conn *db.DBConn) (repository.TokenRepository, error) {
	errs := []error{}

	selectTokenByUserIDPstmt, err := conn.GetPstmt(selectTokenByUserIDQuery)
	errs = append(errs, err)

	insertTokenPstmt, err := conn.GetPstmt(insertTokenQuery)
	errs = append(errs, err)

	deleteTokenByUserIDPstmt, err := conn.GetPstmt(deleteTokenByUserIDQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrapf(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &tokenRepository{
		selectTokenByUserIDPstmt: selectTokenByUserIDPstmt,
		insertTokenPstmt:         insertTokenPstmt,
		deleteTokenByUserIDPstmt: deleteTokenByUserIDPstmt,
	}, nil
}

type tokenRepository struct {
	selectTokenByUserIDPstmt *sql.Stmt
	insertTokenPstmt         *sql.Stmt
	deleteTokenByUserIDPstmt *sql.Stmt
}

func (tR *tokenRepository) FindLatestTokenByUserID(userID string) (*model.Token, *apierror.Error) {
	var token model.Token

	if err := tR.selectTokenByUserIDPstmt.QueryRow(userID).Scan(&token.Token, &token.UserID, &token.CreatedAt, &token.ExpiredAt); err != nil {
		return nil, apierror.NewError(http.StatusNotFound, errors.Wrap(err, "トークンが見つかりません。"))
	}
	return &token, nil
}

func (tR *tokenRepository) StoreToken(t *model.Token) *apierror.Error {
	_, err := tR.insertTokenPstmt.Exec(t.Token, t.UserID, t.CreatedAt, t.ExpiredAt)
	if err != nil {
		return apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "トークンの保存に失敗しました。"))
	}
	return nil
}

func (tR *tokenRepository) RemoveTokenByUserID(userID string) *apierror.Error {
	res, err := tR.deleteTokenByUserIDPstmt.Exec(userID)
	if err != nil {
		return apierror.NewError(http.StatusNotFound, errors.Wrap(err, "トークンの削除に失敗しました。"))
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
