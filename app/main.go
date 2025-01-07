package main

import (
	"fmt"
	"net/http"
)

// Custom handler that wraps ServeMux to handle 404 responses
type customMux struct {
	mux *http.ServeMux
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// List of valid paths
	validPaths := map[string]bool{
		"/":          true,
		"/ascii-art": true,
		//"/download-ascii-art": true, //curl -X POST -d "text=text&banner=standard" -i http://localhost:8080/download-ascii-art
	}

	// Check if the requested path is valid
	if !validPaths[r.URL.Path] {
		http.NotFound(w, r)                           // Return 404 if path is not valid
		fmt.Printf("404 Not Found: %s\n", r.URL.Path) // Optional: Log the unmatched path for debugging
		return
	}

	// Serve the request if the path is valid
	c.mux.ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()

	// Register the main page handler for the root path
	mux.HandleFunc("/", mainPage)

	// Register the /download-ascii-art handler for the export
	mux.HandleFunc("/download-ascii-art", downloadAsciiArtHandler)

	// Register the /ascii-art handler for POST requests
	mux.HandleFunc("/ascii-art", asciiArtHandler)

	// Wrap the mux in our customMux to handle 404 errors
	wrappedMux := &customMux{mux: mux}

	port := 8080
	fmt.Print("Server is running on http://localhost:8080 ")
	err := http.ListenAndServe(":8080", wrappedMux)
	for err != nil {
		port++
		fmt.Printf("\rServer is running on http://localhost:%v ", port)
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), wrappedMux)
	}

}
