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
		log.Fatalf("Error loading .env file")
		return
	}
	cfg := apiConfig{
		APIKey: os.Getenv("API_KEY"),
	}
	const port = "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/power-outages-time", cfg.handlerPowerOutagesTime)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server is running on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
