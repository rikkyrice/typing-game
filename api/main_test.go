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
	"time"

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

	userID := "user02"
	mail := "user01@gmail.com"
	password := "user01"
	createdAt := time.Now().Format("2006-01-02T15:04:05+09:00")
	jsonBlob := []byte(fmt.Sprintf(`{
					"id": "%s",
				 	"mail": "%s",
				 	"password": "%s",
				 	"createdAt": "%s"
				}`, userID, mail, password, createdAt))
	req := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer(jsonBlob))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.Router.ServeHTTP(rec, req)
	defer rg.UserR.RemoveUserByID(userID)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Testユーザー登録ID重複(t *testing.T) {
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

	userID := "rikky"
	mail := "user01@gmail.com"
	password := "user01"
	createdAt := time.Now().Format("2006-01-02T15:04:05+09:00")
	jsonBlob := []byte(fmt.Sprintf(`{
					"id": "%s",
				 	"mail": "%s",
				 	"password": "%s",
				 	"createdAt": "%s"
				}`, userID, mail, password, createdAt))
	req := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer(jsonBlob))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.Router.ServeHTTP(rec, req)
	defer rg.UserR.RemoveUserByID(userID)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
