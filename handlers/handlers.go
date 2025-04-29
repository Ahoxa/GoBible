package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ahoxa/GoBible/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(reqArticle)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	log.Printf("Page: %d\n", page)
	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	// if err != nil {
	// 	http.Error(w, "Invalid article ID", http.StatusBadRequest)
	// 	return
	// }

	article := models.Article1
	json.NewEncoder(w).Encode(article)

}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(reqArticle)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(reqComment)
}
