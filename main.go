package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	APIKey string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	cfg := apiConfig{
		APIKey: os.Getenv("API_KEY"),
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is is not set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/power-outages-time", cfg.handlerPowerOutagesTime)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server is running on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
