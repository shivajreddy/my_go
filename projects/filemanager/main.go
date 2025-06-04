package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	// Check if file exists and create unique name if needed
	originalPath := filePath
	counter := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break // File doesn't exist, we can use this name
		}
		// File exists, try with suffix
		ext := filepath.Ext(originalPath)
		nameWithoutExt := strings.TrimSuffix(filepath.Base(originalPath), ext)
		dir := filepath.Dir(originalPath)
		filePath = filepath.Join(dir, fmt.Sprintf("%s (%d)%s", nameWithoutExt, counter, ext))
		counter++
	}

	// Update fileName to reflect the actual name used
	fileName = filepath.Base(filePath)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to destination
	_, err = io.Copy(dst, file)
	if err != nil {
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
