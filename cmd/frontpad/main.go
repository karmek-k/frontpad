package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/karmek-k/frontpad/pkg/router"
)

func main() {
	if godotenv.Load() != nil {
		panic("no .env file found")
	}

	r := router.CreateRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, r)
}