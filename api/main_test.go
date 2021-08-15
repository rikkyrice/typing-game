package main

import (
	"api/db"
	"api/internal/config"
	"api/internal/domain/model"
	"api/internal/interfaces/router"
	"api/internal/registry"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type wordlistResponse struct {
	WordList *model.WordList `json:"wordlist"`
}

type wordlistsResponse struct {
	Matched   int               `json:"matched"`
	WordLists []*model.WordList `json:"wordlists"`
}

type wordResponse struct {
	Word *model.Word `json:"word"`
}

type wordsQueryRequest struct {
	Words []*model.Word `json:"words"`
}

type wordsResponse struct {
	Matched int           `json:"matched"`
	Words   []*model.Word `json:"words"`
}

type scoreResponse struct {
	Score *model.Score `json:"score"`
}

type scoresResponse struct {
	Matched int            `json:"matched"`
	Scores  []*model.Score `json:"scores"`
}

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
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	userID := "user01"
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
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	userID := "riku"
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

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func Testログイン(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ユーザー登録
	userID := "user01"
	mail := "user01@gmail.com"
	password := "user01"
	createdAt := time.Now().Format("2006-01-02T15:04:05+09:00")
	jsonBlob1 := []byte(fmt.Sprintf(`{
					"id": "%s",
				 	"mail": "%s",
				 	"password": "%s",
				 	"createdAt": "%s"
				}`, userID, mail, password, createdAt))
	req1 := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer(jsonBlob1))
	req1.Header.Set("Content-Type", "application/json")
	rec1 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec1, req1)
	defer rg.UserR.RemoveUserByID(userID)

	assert.Equal(t, http.StatusCreated, rec1.Code)

	time.Sleep(time.Second * 1)
	jsonBlob2 := []byte(fmt.Sprintf(`{
			"password": "%s"
		}`, password))
	req2 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob2))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusOK, rec2.Code)
}

func Testログイン時ユーザーが見つからない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ユーザー登録
	userID := "user01"
	mail := "user01@gmail.com"
	password := "user01"
	createdAt := time.Now().Format("2006-01-02T15:04:05+09:00")
	jsonBlob1 := []byte(fmt.Sprintf(`{
					"id": "%s",
				 	"mail": "%s",
				 	"password": "%s",
				 	"createdAt": "%s"
				}`, userID, mail, password, createdAt))
	req1 := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer(jsonBlob1))
	req1.Header.Set("Content-Type", "application/json")
	rec1 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec1, req1)
	defer rg.UserR.RemoveUserByID(userID)

	assert.Equal(t, http.StatusCreated, rec1.Code)

	time.Sleep(time.Millisecond * 1)

	notfounduser := "notfounduser"
	jsonBlob2 := []byte(fmt.Sprintf(`{
			"password": "%s"
		}`, password))
	req2 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob2))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", notfounduser)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Testログイン時パスワード違い(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ユーザー登録
	userID := "user01"
	mail := "user01@gmail.com"
	password := "user01"
	createdAt := time.Now().Format("2006-01-02T15:04:05+09:00")
	jsonBlob1 := []byte(fmt.Sprintf(`{
					"id": "%s",
				 	"mail": "%s",
				 	"password": "%s",
				 	"createdAt": "%s"
				}`, userID, mail, password, createdAt))
	req1 := httptest.NewRequest(http.MethodPost, "/api/user/signup", bytes.NewBuffer(jsonBlob1))
	req1.Header.Set("Content-Type", "application/json")
	rec1 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec1, req1)
	defer rg.UserR.RemoveUserByID(userID)

	assert.Equal(t, http.StatusCreated, rec1.Code)

	time.Sleep(time.Millisecond * 1)

	wrongPassword := "wrongpassword"
	jsonBlob2 := []byte(fmt.Sprintf(`{
			"password": "%s"
		}`, wrongPassword))
	req2 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob2))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusBadRequest, rec2.Code)
}

func Test単語帳作成(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := &model.WordList{
		ID:          "",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	wlJSON, err := json.Marshal(wl)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPost, "/api/wordlist", bytes.NewBuffer(wlJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wlRes wordlistResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wlRes); err != nil {
		t.FailNow()
	}
	defer rg.WordListR.RemoveWordListByID(wlRes.WordList.ID)

	assert.Equal(t, http.StatusCreated, rec2.Code)
}

func Test単語帳取得(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	req2 := httptest.NewRequest(http.MethodGet, "/api/wordlist", nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wlsRes wordlistsResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wlsRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, 3, wlsRes.Matched)
	assert.Equal(t, wl.ID, wlsRes.WordLists[len(wlsRes.WordLists)-1].ID)
}

func Test単語帳更新(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// 単語帳編集
	wl.Title = "OpenShift"
	wl.Explanation = "OpenShift勉強用の単語帳です。"
	wl.UpdatedAt = time.Now()
	wlJSON, err := json.Marshal(wl)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPut, "/api/wordlist/"+wl.ID, bytes.NewBuffer(wlJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wlRes wordlistResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wlRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusCreated, rec2.Code)
	assert.Equal(t, wl.ID, wlRes.WordList.ID)
	assert.Equal(t, wl.UserID, wlRes.WordList.UserID)
	assert.Equal(t, wl.Title, wlRes.WordList.Title)
	assert.Equal(t, wl.Explanation, wlRes.WordList.Explanation)
	assert.Equal(t, wl.CreatedAt.Format(time.RFC3339), wlRes.WordList.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, wl.UpdatedAt.Format(time.RFC3339), wlRes.WordList.UpdatedAt.Format(time.RFC3339))
}

func Test単語帳更新ID不一致(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	userID := "riku"
	password := "riku"
	jsonBlob1 := []byte(fmt.Sprintf(`{
			"password": "%s"
		}`, password))
	req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("X-User-ID", userID)
	rec1 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec1, req1)

	assert.Equal(t, http.StatusOK, rec1.Code)

	time.Sleep(time.Millisecond * 1)

	var tokenRes tokenResponse
	body, err := ioutil.ReadAll(rec1.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		t.FailNow()
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// 単語帳編集
	wl.Title = "OpenShift"
	wl.Explanation = "OpenShift勉強用の単語帳です。"
	wl.UpdatedAt = time.Now()
	wlJSON, err := json.Marshal(wl)
	if err != nil {
		t.FailNow()
	}
	wrongID := "wrongID"
	req2 := httptest.NewRequest(http.MethodPut, "/api/wordlist/"+wrongID, bytes.NewBuffer(wlJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Test単語帳削除(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	req2 := httptest.NewRequest(http.MethodDelete, "/api/wordlist/"+wl.ID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	_, err = rg.WordListR.FindWordListByID(wl.ID)
	assert.EqualError(t, err, "単語帳が見つかりません。: sql: no rows in result set")
}

func Test単語帳削除ID不一致(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	wrongID := "wrongID"
	req2 := httptest.NewRequest(http.MethodDelete, "/api/wordlist/"+wrongID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Test単語作成(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := &model.Word{
		ID:          "",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "避けられない",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	wJSON, err := json.Marshal(w)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPost, "/api/word", bytes.NewBuffer(wJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wRes wordResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wRes); err != nil {
		t.FailNow()
	}
	defer rg.WordR.RemoveWordByID(wRes.Word.ID)

	assert.Equal(t, http.StatusCreated, rec2.Code)
	assert.Equal(t, w.WordListID, wRes.Word.WordListID)
	assert.Equal(t, w.Word, wRes.Word.Word)
	assert.Equal(t, w.Meaning, wRes.Word.Meaning)
	assert.Equal(t, w.Explanation, wRes.Word.Explanation)
	assert.Equal(t, w.CreatedAt.Format(time.RFC3339), wRes.Word.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, w.UpdatedAt.Format(time.RFC3339), wRes.Word.UpdatedAt.Format(time.RFC3339))
}

func Test単語作成単語帳IDが存在しない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := &model.Word{
		ID:          "",
		WordListID:  "notExists",
		Word:        "inevitable",
		Meaning:     "避けられない",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	wJSON, err := json.Marshal(w)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPost, "/api/word", bytes.NewBuffer(wJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusInternalServerError, rec2.Code)
}

func Test単語作成複数(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語のリスト作成
	ws := []*model.Word{
		{
			ID:          "",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "monster",
			Meaning:     "モンスター",
			Explanation: "Unleash Your Monster",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "energy",
			Meaning:     "エナジー",
			Explanation: "Inject inside your body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	wsq := &wordsQueryRequest{
		Words: ws,
	}
	wsJSON, err := json.Marshal(wsq)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPost, "/api/word/wordlist", bytes.NewBuffer(wsJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wsRes wordsResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wsRes); err != nil {
		t.FailNow()
	}
	defer rg.WordR.RemoveWordByID(wsRes.Words[len(wsRes.Words)-1].ID)
	defer rg.WordR.RemoveWordByID(wsRes.Words[len(wsRes.Words)-2].ID)

	assert.Equal(t, http.StatusCreated, rec2.Code)
	for i := 0; i < 2; i++ {
		assert.Equal(t, ws[i].WordListID, wsRes.Words[len(wsRes.Words)-2+i].WordListID)
		assert.Equal(t, ws[i].Word, wsRes.Words[len(wsRes.Words)-2+i].Word)
		assert.Equal(t, ws[i].Meaning, wsRes.Words[len(wsRes.Words)-2+i].Meaning)
		assert.Equal(t, ws[i].Explanation, wsRes.Words[len(wsRes.Words)-2+i].Explanation)
		assert.Equal(t, ws[i].CreatedAt.Format(time.RFC3339), wsRes.Words[len(wsRes.Words)-2+i].CreatedAt.Format(time.RFC3339))
		assert.Equal(t, ws[i].UpdatedAt.Format(time.RFC3339), wsRes.Words[len(wsRes.Words)-2+i].UpdatedAt.Format(time.RFC3339))
	}
}

func Test単語取得(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	req2 := httptest.NewRequest(http.MethodGet, "/api/word/"+w.ID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wRes wordResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, w.ID, wRes.Word.ID)
	assert.Equal(t, w.WordListID, wRes.Word.WordListID)
	assert.Equal(t, w.Word, wRes.Word.Word)
	assert.Equal(t, w.Meaning, wRes.Word.Meaning)
	assert.Equal(t, w.Explanation, wRes.Word.Explanation)
	assert.Equal(t, w.CreatedAt.Format(time.RFC3339), wRes.Word.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, w.UpdatedAt.Format(time.RFC3339), wRes.Word.UpdatedAt.Format(time.RFC3339))
}

func Test単語取得ID不一致(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	wrongID := "wrongID"
	req2 := httptest.NewRequest(http.MethodGet, "/api/word/wordlist"+wrongID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Test単語取得複数(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// 単語のリスト作成
	ws := []model.Word{
		{
			ID:          "word18ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "bb8938cb-0289-4ed4-9b5e-408d309739ad",
			Word:        "monster",
			Meaning:     "モンスター",
			Explanation: "Unleash Your Monster",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "word28ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "bb8938cb-0289-4ed4-9b5e-408d309739ad",
			Word:        "energy",
			Meaning:     "エナジー",
			Explanation: "Inject inside your body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, w := range ws {
		rg.WordR.CreateWord(w)
		defer rg.WordR.RemoveWordByID(w.ID)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/api/word/wordlist/"+ws[0].WordListID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wsRes wordsResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wsRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, 2, wsRes.Matched)
	idList := []string{ws[0].WordListID, ws[1].WordListID}
	for i := 0; i < 2; i++ {
		assert.Contains(t, idList, wsRes.Words[len(wsRes.Words)-2+i].WordListID)
	}
}

func Test単語取得複数単語帳IDが存在しない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語のリスト作成
	ws := []model.Word{
		{
			ID:          "word18ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "monster",
			Meaning:     "モンスター",
			Explanation: "Unleash Your Monster",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "word28ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "energy",
			Meaning:     "エナジー",
			Explanation: "Inject inside your body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, w := range ws {
		rg.WordR.CreateWord(w)
		defer rg.WordR.RemoveWordByID(w.ID)
	}

	wrongWLID := "wrongwlid"
	req2 := httptest.NewRequest(http.MethodGet, "/api/word/wordlist/"+wrongWLID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wsRes wordsResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wsRes); err != nil {
		t.FailNow()
	}
	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, 0, wsRes.Matched)
}

func Test単語更新(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	// 単語編集
	w.Word = "avoid"
	w.Meaning = "avoid means opposition of inevitable"
	w.Explanation = "you can avoid everything before bad situation."
	w.UpdatedAt = time.Now()
	wJSON, err := json.Marshal(w)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPut, "/api/word/"+w.ID, bytes.NewBuffer(wJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var wRes wordResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &wRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusCreated, rec2.Code)
	assert.Equal(t, w.ID, wRes.Word.ID)
	assert.Equal(t, w.WordListID, wRes.Word.WordListID)
	assert.Equal(t, w.Word, wRes.Word.Word)
	assert.Equal(t, w.Meaning, wRes.Word.Meaning)
	assert.Equal(t, w.Explanation, wRes.Word.Explanation)
	assert.Equal(t, w.CreatedAt.Format(time.RFC3339), wRes.Word.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, w.UpdatedAt.Format(time.RFC3339), wRes.Word.UpdatedAt.Format(time.RFC3339))
}

func Test単語更新ID不一致(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	// 単語編集
	w.Word = "avoid"
	w.Meaning = "avoid means opposition of inevitable"
	w.Explanation = "you can avoid everything before bad situation."
	w.UpdatedAt = time.Now()
	wJSON, err := json.Marshal(w)
	if err != nil {
		t.FailNow()
	}

	wrongID := "wrongID"
	req2 := httptest.NewRequest(http.MethodPut, "/api/word/"+wrongID, bytes.NewBuffer(wJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Test単語削除(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	req2 := httptest.NewRequest(http.MethodDelete, "/api/word/"+w.ID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	_, err = rg.WordR.FindWordByID(w.ID)
	assert.EqualError(t, err, "ID[word38ca-0289-4ed4-9b5e-408d309739ad]の単語が見つかりません。: sql: no rows in result set")
}

func Test単語削除ID不一致(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	wrongID := "wrongid"
	req2 := httptest.NewRequest(http.MethodDelete, "/api/word/"+wrongID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Test単語削除単語帳の削除(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語のリスト作成
	ws := []model.Word{
		{
			ID:          "word18ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "monster",
			Meaning:     "モンスター",
			Explanation: "Unleash Your Monster",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "word28ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "energy",
			Meaning:     "エナジー",
			Explanation: "Inject inside your body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, w := range ws {
		rg.WordR.CreateWord(w)
		defer rg.WordR.RemoveWordByID(w.ID)
	}

	req2 := httptest.NewRequest(http.MethodDelete, "/api/word/wordlist/"+ws[0].WordListID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	words, err := rg.WordR.FindWordByWordListID(ws[0].WordListID)
	assert.Equal(t, 0, len(words))
}

func Test単語削除単語帳IDが見つからない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語のリスト作成
	ws := []model.Word{
		{
			ID:          "word18ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "monster",
			Meaning:     "モンスター",
			Explanation: "Unleash Your Monster",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "word28ca-0289-4ed4-9b5e-408d309739ad",
			WordListID:  "5f52039d-d983-4ebd-90b2-e3e04f821896",
			Word:        "energy",
			Meaning:     "エナジー",
			Explanation: "Inject inside your body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, w := range ws {
		rg.WordR.CreateWord(w)
		defer rg.WordR.RemoveWordByID(w.ID)
	}

	wrongID := "wrongid"
	req2 := httptest.NewRequest(http.MethodDelete, "/api/word/wordlist/"+wrongID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	words, err := rg.WordR.FindWordByWordListID(ws[0].WordListID)
	assert.Equal(t, 2, len(words))
}

func Testスコア作成(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// スコア作成
	s := &model.Score{
		ID:             "",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	sJSON, err := json.Marshal(s)
	if err != nil {
		t.FailNow()
	}
	req2 := httptest.NewRequest(http.MethodPost, "/api/score", bytes.NewBuffer(sJSON))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var sRes scoreResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &sRes); err != nil {
		t.FailNow()
	}
	defer rg.ScoreR.RemoveScoreByID(sRes.Score.ID)

	assert.Equal(t, http.StatusCreated, rec2.Code)
}

func Testスコア取得(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	req2 := httptest.NewRequest(http.MethodGet, "/api/score/"+s.WordListID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var ssRes scoresResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &ssRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, 2, ssRes.Matched)
	assert.Equal(t, s.ID, ssRes.Scores[len(ssRes.Scores)-1].ID)
	assert.Equal(t, s.WordListID, ssRes.Scores[len(ssRes.Scores)-1].WordListID)
	assert.Equal(t, s.ClearTypeCount, ssRes.Scores[len(ssRes.Scores)-1].ClearTypeCount)
	assert.Equal(t, s.MissTypeCount, ssRes.Scores[len(ssRes.Scores)-1].MissTypeCount)
	assert.Equal(t, s.PlayedAt.Format(time.RFC3339), ssRes.Scores[len(ssRes.Scores)-1].PlayedAt.Format(time.RFC3339))
}

func Testスコア取得単語帳IDが存在しない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	wrongWLID := "wrongwlid"
	req2 := httptest.NewRequest(http.MethodGet, "/api/score/"+wrongWLID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var ssRes scoresResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &ssRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, 0, ssRes.Matched)
}

func Test最新スコア取得(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	req2 := httptest.NewRequest(http.MethodGet, "/api/score/"+s.WordListID+"/latest", nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)
	var sRes scoreResponse
	body, err := ioutil.ReadAll(rec2.Body)
	if err != nil {
		t.FailNow()
	}
	if err := json.Unmarshal(body, &sRes); err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, s.ID, sRes.Score.ID)
	assert.Equal(t, s.WordListID, sRes.Score.WordListID)
	assert.Equal(t, s.ClearTypeCount, sRes.Score.ClearTypeCount)
	assert.Equal(t, s.MissTypeCount, sRes.Score.MissTypeCount)
	assert.Equal(t, s.PlayedAt.Format(time.RFC3339), sRes.Score.PlayedAt.Format(time.RFC3339))
}

func Test最新スコア取得単語帳IDが存在しない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	wrongWLID := "wrongwlid"
	req2 := httptest.NewRequest(http.MethodGet, "/api/score/"+wrongWLID+"/latest", nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNotFound, rec2.Code)
}

func Testスコア削除(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	req2 := httptest.NewRequest(http.MethodDelete, "/api/score/"+s.WordListID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	_, err = rg.ScoreR.FIndLatestScoreByWordListID(s.WordListID)
	assert.EqualError(t, err, "ID[bb8938cb-0289-4ed4-9b5e-408d309739ad]の単語帳の最新のスコアが見つかりません。: sql: no rows in result set")
}

func Testスコア削除単語帳IDが存在しない(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	wrongID := "wrongid"
	req2 := httptest.NewRequest(http.MethodDelete, "/api/score/"+wrongID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	scores, err := rg.ScoreR.FindScoreByWordListID(s.WordListID)
	assert.Equal(t, 1, len(scores))
}

// Test単語がある状態の単語帳を削除
func Test単語がある状態の単語帳を削除(t *testing.T) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
	rg, _ := registry.NewRegistry(conn)
	r := router.NewRouter()
	r.Init(rg)

	// ログイン
	var tokenRes tokenResponse
	userID := "riku"
	isLoggedIn := true
	token, apierr := rg.TokenR.FindLatestTokenByUserID(userID)
	if apierr != nil {
		isLoggedIn = false
	}
	if token != nil {
		valid := time.Now()
		if valid.After(token.ExpiredAt) {
			isLoggedIn = false
		}
	}

	if !isLoggedIn {
		password := "riku"
		jsonBlob1 := []byte(fmt.Sprintf(`{
				"password": "%s"
			}`, password))
		req1 := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonBlob1))
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("X-User-ID", userID)
		rec1 := httptest.NewRecorder()

		r.Router.ServeHTTP(rec1, req1)

		assert.Equal(t, http.StatusOK, rec1.Code)

		body, err := ioutil.ReadAll(rec1.Body)
		if err != nil {
			t.FailNow()
		}
		if err := json.Unmarshal(body, &tokenRes); err != nil {
			t.FailNow()
		}
	} else {
		tokenRes.Token = token.Token
	}

	// 単語帳作成
	wl := model.WordList{
		ID:          "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		UserID:      userID,
		Title:       "TOEIC",
		Explanation: "TOEIC勉強用の単語帳です。",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordListR.CreateWordList(wl)
	defer rg.WordListR.RemoveWordListByID(wl.ID)

	// 単語作成
	w := model.Word{
		ID:          "word38ca-0289-4ed4-9b5e-408d309739ad",
		WordListID:  "bb8938cb-0289-4ed4-9b5e-408d309739ad",
		Word:        "inevitable",
		Meaning:     "inevitable means you cannot avoid",
		Explanation: "Your undoing is now inevitable.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rg.WordR.CreateWord(w)
	defer rg.WordR.RemoveWordByID(w.ID)

	// スコア作成
	s := model.Score{
		ID:             "score28c-0289-4ed4-9b5e-408d309739ad",
		WordListID:     "5f52039d-d983-4ebd-90b2-e3e04f821896",
		ClearTypeCount: 140,
		MissTypeCount:  40,
		PlayedAt:       time.Now(),
	}
	rg.ScoreR.CreateScore(s)
	defer rg.ScoreR.RemoveScoreByID(s.ID)

	req2 := httptest.NewRequest(http.MethodDelete, "/api/wordlist/"+wl.ID, nil)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("X-User-ID", userID)
	req2.Header.Set("Authorization", "Bearer "+tokenRes.Token)
	rec2 := httptest.NewRecorder()

	r.Router.ServeHTTP(rec2, req2)

	assert.Equal(t, http.StatusNoContent, rec2.Code)

	_, err = rg.WordListR.FindWordListByID(wl.ID)
	assert.EqualError(t, err, "単語帳が見つかりません。: sql: no rows in result set")

	_, err = rg.WordR.FindWordByID(w.ID)
	assert.EqualError(t, err, "ID[word38ca-0289-4ed4-9b5e-408d309739ad]の単語が見つかりません。: sql: no rows in result set")

	_, err = rg.ScoreR.FIndLatestScoreByWordListID(wl.ID)
	assert.EqualError(t, err, "ID[bb8938cb-0289-4ed4-9b5e-408d309739ad]の単語帳の最新のスコアが見つかりません。: sql: no rows in result set")
}
