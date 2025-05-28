package main

import (
	"log"
	"net/http"
	"newsgo/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/result", controllers.Results)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
