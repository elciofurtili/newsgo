package controllers

import (
	"html/template"
	"net/http"
	"newsgo/models"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.gohtml", "views/results.gohtml"))
	query := r.URL.Query().Get("q")
	var articles []models.Article

	if query != "" {
		apiKey := os.Getenv("NEWSAPI_KEY")
		articles = models.FetchNews(query, apiKey)
	}

	data := struct {
		Query    string
		Articles []models.Article
	}{
		Query:    query,
		Articles: articles,
	}

	tpl.Execute(w, data)
}

func Results(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/results.gohtml"))
	query := r.URL.Query().Get("q")
	var articles []models.Article

	if query != "" {
		apiKey := os.Getenv("NEWSAPI_KEY")
		articles = models.FetchNews(query, apiKey)
	}

	data := struct {
		Query    string
		Articles []models.Article
	}{
		Query:    query,
		Articles: articles,
	}

	tpl.ExecuteTemplate(w, "results", data)
}
