package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("FILE MANAGER")

	// Serve the main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle file uploads
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form (32MB max)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the file from form data
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Use hardcoded desktop path
	desktopPath := "/mnt/c/Users/sreddy/Desktop"

	// Create the full file path
	fileName := header.Filename
	filePath := filepath.Join(desktopPath, fileName)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to destination
	totalBytesWritten, copyErr := io.Copy(dst, file)
	fmt.Println("writtenResult", totalBytesWritten)
	if copyErr != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<html>
		<body>
			<h2>File uploaded successfully!</h2>
			<p>File "%s" saved to desktop</p>
			<a href="/">Upload another file</a>
		</body>
		</html>
	`, fileName)
}
