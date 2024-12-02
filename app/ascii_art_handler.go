package main

import (
	"net/http"
)

// ArtData holds data to be passed to templates (input, generated ASCII art, banner, errors)
type ArtData struct {
	InputText string
	ArtResult string
	Banner    string
	Error     string
}

// asciiArtHandler processes the ASCII art request and displays the result page
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request - Unable to parse form data", http.StatusBadRequest)
		return
	}

	inputText := r.FormValue("text")
	alignment := r.FormValue("alignment")
	color := r.FormValue("color")
	banner := r.FormValue("banner")

	// Check if the text is not empty
	if inputText == "" {
		http.Error(w, "type something", http.StatusBadRequest)
		return
	}

	// Check if the alignment is valid
	if alignment != "left" && alignment != "center" && alignment != "right" {
		http.Error(w, "Bad Request - Invalid alignment option", http.StatusBadRequest)
		return
	}

	// Check if the color is valid
	if color != "#ffffff" && color != "#ff0000" && color != "#00ff00" && color != "#0000ff" && color != "#ffff00" && color != "#ff00ff" && color != "#00ffff" && color != "#000000" {
		http.Error(w, "Bad Request - Invalid color option", http.StatusBadRequest)
		return
	}

	// Check if the provided banner is valid
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Bad Request - Invalid banner template", http.StatusBadRequest)
		return
	}

	artResult, err := GenerateAsciiArt(inputText, banner)
	if err != nil {
		if err.Error() == "banner file not found" {
			http.Error(w, "Not Found - Banner file missing", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Set content type to text/plain to ensure proper display
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(artResult))
}
