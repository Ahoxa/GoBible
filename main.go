package main

import (
	"log"
	"net/http"

	"github.com/ahoxa/GoBible/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods("POST")
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods("GET")
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleDetailHandler).Methods("GET")
	r.HandleFunc("article/nice", handlers.PostNiceHandler).Methods("POST")
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
