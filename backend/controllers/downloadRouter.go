package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// HandleDownloadZip is the handler for generating and downloading the zip file
func HandleDownloadZip(w http.ResponseWriter, r *http.Request) {
	// The path where uploaded files are stored
	uploadDir := "./uploads" // Change this path according to your project structure

	// Generate a new zip file
	zipFileName := "project.zip"
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create zip file: %v", err), http.StatusInternalServerError)
		return
	}
	defer zipFile.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the uploaded files directory and add files to the zip
	err = filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a new zip entry for each file
		zipEntry, err := zipWriter.Create(info.Name())
		if err != nil {
			return err
		}

		// Copy the file content into the zip entry
		_, err = io.Copy(zipEntry, file)
		return err
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding files to zip: %v", err), http.StatusInternalServerError)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+zipFileName)

	// Serve the generated zip file to the client
	http.ServeFile(w, r, zipFileName)
}
