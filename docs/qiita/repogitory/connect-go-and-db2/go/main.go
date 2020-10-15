package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
	"local.packages/go/model"
)

func main() {
	config := "HOSTNAME=localhost;DATABASE=USERDB;PORT=50000;UID=db2inst1;PWD=password"
	conn, err := sql.Open("go_ibm_db", config)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()

	users, err := model.GetAllUser(conn)
	if err != nil {
		fmt.Printf("取得に失敗 %+v", err)
	}

	// 件数少ないので3重for文で。
	for _, user := range users {
		fmt.Println(user.String())
		tweets, err := model.GetAllTweets(conn, user.GetID())
		if err != nil {
			fmt.Printf("取得に失敗 %+v", err)
		}
		for _, tweet := range tweets {
			fmt.Println(tweet.String())
			replys, err := model.GetAllReplys(conn, tweet.GetID())
			if err != nil {
				fmt.Printf("取得に失敗", err)
			}
			for _, reply := range replys {
				fmt.Println(reply.String())
				fmt.Println(reply.GetBody())
			}
		}
		fmt.Println("-----------------------")
	}
}
