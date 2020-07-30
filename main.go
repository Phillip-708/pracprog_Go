package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var hashmap map[int]string

func main() {
	hashmap = make(map[int]string)
	http.HandleFunc("/todo", todo)
	http.ListenAndServe(":8080", nil)
}

func todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "todoを表示します\n")
		fmt.Fprint(w, hashmap)
	case http.MethodPost:
		bufbody := new(bytes.Buffer)
		_, err := bufbody.ReadFrom(r.Body)
		if err != nil {
			log.Println(err)
		}
		body := bufbody.String()
		hashmap[len(hashmap)] = body
		fmt.Fprint(w, "todoを追加します\n")
	case http.MethodPut:
		fmt.Fprint(w, "todoを更新します\n")
	case http.MethodDelete:
		fmt.Fprint(w, "todoを削除します\n")
	default:
		fmt.Fprint(w, "エラー \n")
	}

}
