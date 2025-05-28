package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func FetchNews(query, apiKey string) []Article {
	url := "https://newsapi.org/v2/everything?q=" + query +
		"&sortBy=popularity&pageSize=10&language=pt&apiKey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Erro na requisição:", err)
		return nil
	}
	defer resp.Body.Close()

	var result NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Erro ao decodificar:", err)
		return nil
	}

	return result.Articles
}
