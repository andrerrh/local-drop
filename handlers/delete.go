package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid delete request", http.StatusMethodNotAllowed)
		return
	}

	files := r.URL.Query()["files"]
	if len(files) == 0 {
		http.Error(w, "No files specified for deletion", http.StatusBadRequest)
		return
	}

	for _, fileName := range files {
		filePath := filepath.Join("files", fileName)

		err := os.Remove(filePath)
		if err != nil {
			http.Error(w, "Failed to delete file(s)", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ``)
}
