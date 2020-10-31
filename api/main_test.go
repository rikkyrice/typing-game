package main

import (
	"api/db"
	"api/internal/config"
	"api/internal/interfaces/router"
	"api/internal/registry"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testユーザー登録(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	jsonStr := `{"id": "user01",
				 "mail": "user01@user.com",
				 "password": "p@ssword",
				 "createdAt": "2020-08-05T00:00:00+09:00"
				}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.Router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
}
