package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"main.go/app/config"
	"main.go/app/models"
	"main.go/app/utils"
)

func GenerateShortCode(w http.ResponseWriter, r *http.Request) {
    // 1. Limit body size to prevent memory exhaustion attacks
    r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit

    var url models.Url
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, "Invalid JSON request", http.StatusBadRequest)
        return
    }

    // 2. Validate URL input
    if url.OriginalUrl == "" {
        http.Error(w, "original_url is required", http.StatusBadRequest)
        return
    }

    code, err := utils.GenerateShortCode(10)
    if err != nil {
        http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
        return
    }

    // 3. Use Exec instead of QueryRow for simple INSERT statements
    query := `INSERT INTO urls (short_code, original_url) VALUES ($1, $2)`
    _, err = config.DB.Exec(query, code, url.OriginalUrl)
    if err != nil {
        http.Error(w, "Failed to save URL to database", http.StatusInternalServerError)
        return
    }

    // 4. Return correct 201 Created status and include the short code
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success":    true,
        "message":    "Generated the short code",
        "short_code": code,
    })
}

func RedirectURL(w http.ResponseWriter,r *http.Request){
    shortCode :=chi.URLParam(r,"code")
    if shortCode == "" {
		http.NotFound(w, r)
		return
	}

    var originalURL string
    query := `SELECT original_url FROM urls WHERE short_code = $1`
    err := config.DB.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
    http.Redirect(w,r, originalURL, http.StatusFound)
}