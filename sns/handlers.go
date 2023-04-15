package main

import (
	"fmt"
	"html/template"
"net/http"
)

// HTTPハンドラを集めた型
type Handler struct {
	commentList *CommentList
}

// Handlersを作成する
func NewHandler(commentList *CommentList) *Handler {
	return &Handler{commentList: commentList}
}
// ListHandlerで仕様するテンプレート
var listTemplate = template.Must(template.ParseFiles("list.html"))

// データを表示するハンドラ
func (handler *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	comments, error := handler.commentList.GetComment()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	if error := listTemplate.Execute(w, comments); error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
}

// 保存
func (handler *Handler) SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method != http.MethodPost {
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
		return
	}

	comment := r.FormValue("comment")
	if comment == "" {
		http.Error(w, "文字が入力されていません。", http.StatusBadRequest)
		return
	}

	serihu := &Comment{
		Comment: comment,
	}

	if error:= handler.commentList.AddComment(serihu); error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)

}

//// 集計を表示するハンドラ
//var fieldTemplate = template.Must(template.ParseFiles("Post.html"))
//
//// 集計を表示するハンドラ
//func (handler *Handler) SummaryHandler(w http.ResponseWriter, r *http.Request){
//	summeries, err := handler.commentList.GetComment()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	if err:= fieldTemplate.Execute(w, summeries); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
