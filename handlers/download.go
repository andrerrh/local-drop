package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var fileName string

	if queryFile := r.URL.Query().Get("file"); queryFile != "" {
		fileName = queryFile
	} else {
		segments := strings.Split(r.URL.Path, "/")
		if len(segments) > 2 {
			fileName = segments[len(segments)-1]
		}
	}

	if fileName == "" {
		http.Error(w, "File parameter is missing", http.StatusBadRequest)
	}

	currFilePath := path.Join("files", fileName)
	if _, err := os.Stat(currFilePath); err != nil {
		fmt.Println("Error retrieving file: " + fileName)
		http.Error(w, "Error retrieving file: "+fileName+"\nMight not exist", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	fmt.Println("Serving file: " + fileName)
	http.ServeFile(w, r, currFilePath)
}
