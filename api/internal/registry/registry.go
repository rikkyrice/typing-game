package registry

import (
	"api/db"

	"api/internal/application/usecase/auth"
	"api/internal/application/usecase/user"
	"api/internal/domain/repository"
	"api/internal/infrastructure"
	"api/internal/interfaces/handler"
)

// Registry 必要な要素を登録する構造体
type Registry struct {
	TokenR repository.TokenRepository
	UserR  repository.UserRepository

	AuthUC       auth.AuthUseCase
	UserSignupUC user.UserSignupUseCase
	UserLoginUC  user.UserLoginUseCase

	HealthCheckH handler.HealthCheckHandler
	UserH        handler.UserHandler
}

// NewRegistry レジストリを生成
func NewRegistry(conn *db.DBConn) *Registry {
	tR, err := infrastructure.NewTokenRepository(conn)
	if err != nil {
		return nil
	}
	uR, err := infrastructure.NewUserRepository(conn)
	if err != nil {
		return nil
	}

	aUC := auth.NewAuthUseCase(tR, uR)

	uSUC := user.NewUserSignupUseCase(uR, aUC)
	uLUC := user.NewUserLoginUseCase(uR, aUC)

	hH := handler.NewHealthCheckHandler()
	uH := handler.NewUserHandler(uSUC, uLUC)

	return &Registry{
		TokenR:       tR,
		UserR:        uR,
		AuthUC:       aUC,
		UserSignupUC: uSUC,
		UserLoginUC:  uLUC,
		HealthCheckH: hH,
		UserH:        uH,
	}
}
