package main

import (
	"database/sql"
	"fmt"
)

type Comment struct {
	Id int
	Comment string
}

// コメントの処理を行う型
type CommentList struct {
	db *sql.DB
}

// 新しいコメントリストを作成
func NewCommentList(db *sql.DB) *CommentList {
	return &CommentList{db: db}
}

// テーブルがなかったら作成する
func (commentList *CommentList) CreateTableComment() error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		comment MESSAGE_TEXT NOT NULL
	);`

	_, err := commentList.db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

//データベースに新しいCommentを追加する
func (commentList *CommentList) AddComment(comments *Comment) error {
	fmt.Println(comments)
	const sqlStr = `INSERT INTO comments(comment) VALUES(?);`
	_, err := commentList.db.Exec(sqlStr, comments.Comment)
	if err != nil {
		return err
	}
	return nil
}

// これまでのメッセージを取得
// エラーが発生したら第２戻り値で返す
func (commentList *CommentList) GetComment() ([]*Comment, error) {
	const sqlStr = `SELECT * FROM comments ORDER BY id DESC`
	rows, err := commentList.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Id, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}