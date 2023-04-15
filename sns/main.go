package main

import (

	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "snsAccount.db")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// snsAccountをNewAccountBookを使って作成
	commentList := NewCommentList(db)

	// テーブルを作成
	if err := commentList.CreateTableComment(); err != nil {
		log.Fatal(err)
	}

	// HandlersをNewHandlersを使って作成
	handler := NewHandler(commentList)

	// ハンドラの登録
	http.HandleFunc("/", handler.ListHandler)
	//http.Handle("/templete/", http.StripPrefix("/templete/", http.FileServer(http.Dir("templete"))))
	http.HandleFunc("/save", handler.SaveHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
