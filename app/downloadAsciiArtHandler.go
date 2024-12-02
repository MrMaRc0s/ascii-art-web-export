package main

import "net/http"

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

	// Generate ASCII art
	artResult, err := GenerateAsciiArt(inputText, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Serve ASCII art as a downloadable file
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(artResult))
}
