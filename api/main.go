package main

import (
	"log"

	"api/db"
	"api/internal/config"
	"api/internal/interfaces/router"
	"api/internal/registry"
)

func main() {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		log.Fatalf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		log.Fatalf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()

	rg, err := registry.NewRegistry(conn)
	if err != nil {
		log.Fatalf("registryの初期化に失敗しました。%+v", err)
	}
	r := router.NewRouter()

	r.Init(rg)

	r.StartServer(":1323")
}
