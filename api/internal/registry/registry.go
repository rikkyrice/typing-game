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
	TokenR    repository.TokenRepository
	UserR     repository.UserRepository
	WordListR repository.WordListRepository
	WordR     repository.WordRepository

	AuthUC     usecase.AuthUseCase
	UserUC     usecase.UserUseCase
	WordListUC usecase.WordListUseCase
	WordUC     usecase.WordUseCase

	HealthCheckH handler.HealthCheckHandler
	UserH        handler.UserHandler
	WordListH    handler.WordListHandler
	WordH        handler.WordHandler
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
	wR, err := infrastructure.NewWordRepository(conn)
	if err != nil {
		return nil
	}

	aUC := usecase.NewAuthUseCase(tR, uR)
	uUC := usecase.NewUserUseCase(uR, aUC)
	wlUC := usecase.NewWordListUseCase(wlR)
	wUC := usecase.NewWordUseCase(wR)

	hH := handler.NewHealthCheckHandler()
	uH := handler.NewUserHandler(uUC)
	wlH := handler.NewWordListHandler(wlUC)
	wH := handler.NewWordHandler(wUC)

	return &Registry{
		TokenR:       tR,
		UserR:        uR,
		WordListR:    wlR,
		WordR:        wR,
		AuthUC:       aUC,
		UserUC:       uUC,
		WordListUC:   wlUC,
		WordUC:       wUC,
		HealthCheckH: hH,
		UserH:        uH,
		WordListH:    wlH,
		WordH:        wH,
	}
}
