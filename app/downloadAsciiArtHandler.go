package main

import (
	"archive/zip"
	"bytes"
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

	// Generate ASCII art
	artResult, err := GenerateAsciiArt(inputText, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a buffer to store the ZIP file in memory
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// Add the ASCII art as a file in the ZIP archive
	zipFile, err := zipWriter.Create("ascii_art.txt")
	if err != nil {
		http.Error(w, "Internal Server Error - Unable to create ZIP file", http.StatusInternalServerError)
		return
	}
	_, err = zipFile.Write([]byte(artResult))
	if err != nil {
		http.Error(w, "Internal Server Error - Unable to write to ZIP file", http.StatusInternalServerError)
		return
	}

	// Close the ZIP writer to finalize the archive
	err = zipWriter.Close()
	if err != nil {
		http.Error(w, "Internal Server Error - Unable to finalize ZIP file", http.StatusInternalServerError)
		return
	}

	// Serve the ZIP file as a downloadable file
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.zip")
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
