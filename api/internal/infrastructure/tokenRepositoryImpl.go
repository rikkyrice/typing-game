package infrastructure

import (
	"api/db"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

const selectTokenByUserIDQuery string = `
	SELECT *
	FROM token
	WHERE user_id = ?
`

const insertTokenQuery string = `
	INSERT INTO token
	VALUES(?,?,?,?)
`

const deleteTokenByUserIDQuery string = `
	DELETE FROM token
	WHERE user_id = ?
`

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

func (tR *tokenRepository) FindTokenByUserID(ctx context.Context, userID string) (model.Token, error) {
	var token model.Token

	if err := tR.selectTokenByUserIDPstmt.QueryRow(userID).Scan(&token.Token, &token.UserID, &token.CreatedAt, &token.ExpiredAt); err != nil {
		return model.Token{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	return token, nil
}

func (tR *tokenRepository) StoreToken(t *model.Token) error {
	_, err := tR.insertTokenPstmt.Exec(t.Token, t.UserID, t.CreatedAt, t.ExpiredAt)
	if err != nil {
		return errors.Wrap(err, "トークンの作成に失敗しました。")
	}
	return err
}

func (tR *tokenRepository) RemoveTokenByUserID(ctx context.Context, userID string) error {
	_, err := tR.deleteTokenByUserIDPstmt.Exec(userID)
	if err != nil {
		return errors.Wrap(err, "トークンの削除に失敗しました。")
	}
	return nil
}
