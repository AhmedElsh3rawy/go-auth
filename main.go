package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /", getHello)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("[server]: running on localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error starting server: %v\n", err)
		return
	}
}
