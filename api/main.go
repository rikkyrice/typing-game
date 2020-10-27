package main

import (
	"fmt"

	"api/db"
	"api/internal/config"
	"api/internal/interfaces/router"
	"api/internal/registry"
)

// TODO: usecaseはサービスにあたるので一つでいい。なんならパッケージに分ける必要もない。

func main() {
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

	r.StartServer(":1323")
}
