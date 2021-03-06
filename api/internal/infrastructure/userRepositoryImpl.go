package infrastructure

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"database/sql"
	"net/http"

	"api/db"

	"github.com/pkg/errors"
)

const selectUserByIDQuery string = `
	SELECT * FROM users WHERE user_id = $1
`

const insertUserQuery string = `
	INSERT INTO users
	VALUES($1,$2,$3,$4)
`

const deleteUserQuery string = `
	DELETE FROM users WHERE user_id = $1
`

// NewUserRepository ユーザーリポジトリの生成
func NewUserRepository(conn *db.DBConn) (repository.UserRepository, error) {
	errs := []error{}

	selectUserByIDPstmt, err := conn.GetPstmt(selectUserByIDQuery)
	errs = append(errs, err)

	insertUserPstmt, err := conn.GetPstmt(insertUserQuery)
	errs = append(errs, err)

	deleteUserPstmt, err := conn.GetPstmt(deleteUserQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrap(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &userRepository{
		selectUserByIDPstmt: selectUserByIDPstmt,
		insertUserPstmt:     insertUserPstmt,
		deleteUserPstmt:     deleteUserPstmt,
	}, nil
}

// userRepository ユーザーのリポジトリインターフェース
type userRepository struct {
	selectUserByIDPstmt *sql.Stmt
	insertUserPstmt     *sql.Stmt
	deleteUserPstmt     *sql.Stmt
}

func (uR *userRepository) FindUserByID(userID string) (*model.User, *apierror.Error) {
	var user model.User

	if err := uR.selectUserByIDPstmt.QueryRow(userID).Scan(&user.ID, &user.Mail, &user.Password, &user.CreatedAt); err != nil {
		return nil, apierror.NewError(http.StatusNotFound, errors.Wrap(err, "ユーザーが見つかりません。"))
	}
	return &user, nil
}

func (uR *userRepository) CreateUser(user model.User) (string, *apierror.Error) {
	_, err := uR.insertUserPstmt.Exec(user.ID, user.Mail, user.Password, user.CreatedAt)
	if err != nil {
		return "", apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "ユーザーの作成に失敗しました。"))
	}
	return user.ID, nil
}

func (uR *userRepository) RemoveUserByID(userID string) *apierror.Error {
	res, err := uR.deleteUserPstmt.Exec(userID)
	if err != nil {
		return apierror.NewError(http.StatusNotFound, errors.Wrap(err, "ユーザーの削除に失敗しました。"))
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
