package registry

import (
	"api/db"

	"api/internal/application/usecase"
	"api/internal/domain/repository"
	"api/internal/infrastructure"
	"api/internal/interfaces/handler"
)

// Registry 必要な要素を登録する構造体
type Registry struct {
	TokenR repository.TokenRepository
	UserR  repository.UserRepository

	AuthUC     usecase.AuthUseCase
	UserUC     usecase.UserUseCase
	WordListUC usecase.WordListUseCase

	HealthCheckH handler.HealthCheckHandler
	UserH        handler.UserHandler
	WordListH    handler.WordListHandler
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
	wlR, err := infrastructure.NewWordListRepository(conn)
	if err != nil {
		return nil
	}

	aUC := usecase.NewAuthUseCase(tR, uR)
	uUC := usecase.NewUserUseCase(uR, aUC)
	wlUC := usecase.NewWordListUseCase(wlR)

	hH := handler.NewHealthCheckHandler()
	uH := handler.NewUserHandler(uUC)
	wlH := handler.NewWordListHandler(wlUC)

	return &Registry{
		TokenR:       tR,
		UserR:        uR,
		AuthUC:       aUC,
		UserUC:       uUC,
		WordListUC:   wlUC,
		HealthCheckH: hH,
		UserH:        uH,
		WordListH:    wlH,
	}
}
