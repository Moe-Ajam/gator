package main

import (
	"database/sql"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	dbUrl := os.Getenv("CONN")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Could not establish database connection", err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
