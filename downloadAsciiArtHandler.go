package main

import (
	"net/http"
	"strconv"
)

func downloadAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request - Unable to parse form data", http.StatusBadRequest)
		return
	}

	inputText := r.FormValue("text")
	banner := r.FormValue("banner")
	if inputText == "" {
		http.Error(w, "Bad Request - Empty Text", http.StatusBadRequest)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Bad Request - Invalid banner template", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	artResult, err := GenerateAsciiArt(inputText, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	artLength := len(artResult)

	// Serve ASCII art as a downloadable file
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(artLength))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(artResult))
}
