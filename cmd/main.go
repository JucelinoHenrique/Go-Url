package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-url-shortener/internal/handler"
	"go-url-shortener/internal/model"
	"go-url-shortener/internal/repository"
)

func main() {

	// Try to connect to the database and migrate the schema
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		// if fails, log the error and exit
		log.Fatal("Failed to connect to database:", err)
	}

	// db.AutoMigrate ensure that the table in DB matches the model structure
	err = db.AutoMigrate(&model.URL{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")

	// Create repository and handler instances
	urlRepo := repository.NewURLRepository(db)
	urlHandler := handler.NewURLHandler(urlRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", urlHandler.ShortenURL)
	mux.HandleFunc("/list", urlHandler.ListURLs)
	mux.HandleFunc("/", urlHandler.RedirectURL)
	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
}
