package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/util"

	"gorm.io/gorm"
)

type URLHandler struct {
	urlRepo *repository.URLRepository
}

func NewURLHandler(urlRepo *repository.URLRepository) *URLHandler {
	return &URLHandler{urlRepo: urlRepo}
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	shortCode := util.GenerateShortCode(6)

	newURL := model.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		Clicks:      0,
	}

	if err := h.repo.Save(&newURL); err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"short_code":   shortCode,
		"short_url":    fmt.Sprintf("http://localhost:8080/%s", shortCode),
		"original_url": req.URL,
	}

	json.NewEncoder(w).Encode(response)

}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	shortCode := r.URL.Path[1:]
	if shortCode == "" {
		fmt.Fprintln(w, "Welcome to the URL Shortener Service! USE /shorten to shorten URLs.")
		return
	}

	url, err := h.urlRepo.GetByShortCode(shortCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error internal server", http.StatusInternalServerError)
	}

	if err := h.urlRepo.IncrementClicks(shortCode); err != nil {
		log.Println("Failed to increment click count:", err)
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
