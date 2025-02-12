package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default if PORT isn't set
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Kira!, I'm glad to serve you")
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
