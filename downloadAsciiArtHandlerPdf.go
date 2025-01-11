package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func downloadAsciiArtHandlerPdf(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request - Unable to parse form data", http.StatusBadRequest)
		return
	}

	input := r.FormValue("textpdf")
	banner := r.FormValue("bannerpdf")
	inputText := ""
	inputLenght := 0
	if input == "" {
		http.Error(w, "Bad Request - Empty Text", http.StatusBadRequest)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Bad Request - Invalid banner template", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(input); i++ {
		inputText += string(input[i])
		inputLenght += checkLenght(string(input[i]), banner)
		if inputLenght >= 680 {
			inputText += "\\n"
			inputLenght = 0
		}
	}

	// Generate ASCII art
	artResult, err := GenerateAsciiArt(inputText, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Build a simple PDF manually
	pdfContent := buildMinimalPDF(artResult)

	// Serve the PDF as a downloadable file
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Length", strconv.Itoa(len(pdfContent)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pdfContent))
}

func buildMinimalPDF(content string) string {
	// Start of the PDF file
	pdf := "%PDF-1.4\n"
	pdf += "1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n"
	pdf += "2 0 obj\n<< /Type /Pages /Count 1 /Kids [3 0 R] >>\nendobj\n"
	pdf += "3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Contents 4 0 R /Resources << /Font << /F1 5 0 R >> >> >>\nendobj\n"

	// Initialize the text content stream
	textStream := "BT\n/F1 10 Tf\n" // Begin text and set font size to 10
	yOffset := 770                  // Start near the top of the page (y-axis)
	for _, line := range splitByLine(content) {
		// Break lines dynamically to avoid overflow
		if yOffset < 20 {
			// Add logic for handling page overflow (not implemented in this basic example)
			break
		}
		// Move to a new position for each line and add the line's text
		textStream += fmt.Sprintf("1 0 0 1 50 %d Tm (%s) Tj\n", yOffset, escapePDFString(line))
		yOffset -= 12 // Adjust spacing between lines
	}
	textStream += "ET\n" // End text

	// Add the content stream object
	pdf += fmt.Sprintf("4 0 obj\n<< /Length %d >>\nstream\n%s\nendstream\nendobj\n", len(textStream), textStream)

	// Add the font object
	pdf += "5 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Courier >>\nendobj\n"

	// Cross-reference table and trailer
	pdf += "xref\n0 6\n0000000000 65535 f \n0000000010 00000 n \n0000000077 00000 n \n0000000178 00000 n \n0000000358 00000 n \n0000000439 00000 n \n"
	pdf += "trailer\n<< /Root 1 0 R /Size 6 >>\nstartxref\n540\n%%EOF"

	return pdf
}

func escapePDFString(s string) string {
	// Escape special characters for PDF strings
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "\\", "\\\\"), "(", "\\("), ")", "\\)")
}

func splitByLine(content string) []string {
	// Split content into lines
	return strings.Split(content, "\n")
}

func checkLenght(r string, banner string) int {
	artResult, _ := GenerateAsciiArt(r, banner)
	return len(artResult)
}
