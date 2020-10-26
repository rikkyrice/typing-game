package infrastructure

import (
	"api/internal/domain/model"
	"database/sql"

	"api/db"

	"github.com/pkg/errors"
)

const selectUserByIDQuery string = `
	SELECT * FROM users WHERE user_id = ?
`

const insertUserQuery string = `
	INSERT INTO users
	VALUES(?,?,?,?)
`

func NewUserRepository(conn *db.DBConn) (*userRepository, error) {
	errs := []error{}

	selectUserByIDPstmt, err := conn.GetPstmt(selectUserByIDQuery)
	errs = append(errs, err)

	insertUserPstmt, err := conn.GetPstmt(insertUserQuery)
	errs = append(errs, err)

	// いずれかのステートメント生成が失敗した場合にはエラーを返す
	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrapf(err, "ステートメントの作成に失敗しました。")
		}
	}

	return &userRepository{
		selectUserByIDPstmt: selectUserByIDPstmt,
		insertUserPstmt:     insertUserPstmt,
	}, nil
}

// UserRepository ユーザーのリポジトリインターフェース
type userRepository struct {
	selectUserByIDPstmt *sql.Stmt
	insertUserPstmt     *sql.Stmt
}

func (uR *userRepository) FindUserByID(userID string) (model.User, error) {
	var user model.User

	if err := uR.selectUserByIDPstmt.QueryRow(userID).Scan(&user.ID, &user.Mail, &user.Password, &user.CreatedAt); err != nil {
		return model.User{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	return user, nil
}

func (uR *userRepository) CreateUser(user model.User) (string, error) {
	_, err := uR.insertUserPstmt.Exec(user.ID, user.Mail, user.Password, user.CreatedAt)
	if err != nil {
		return "", errors.Wrap(err, "ユーザーの作成に失敗しました。")
	}
	return user.ID, err
}
