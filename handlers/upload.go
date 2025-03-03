package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(1 << 30) //1GB Limit
	if err != nil {
		http.Error(w, "File(s) too large", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	var uploadedFiles []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open a specific file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dstPath := filepath.Join("files", fileHeader.Filename)
		dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
		if err != nil {
			if os.IsExist(err) {
				w.WriteHeader(http.StatusConflict) // Set the correct HTTP status code
				fmt.Fprintf(w, `<span remove-me="2s"
					class="bg-red-500
					text-white
					rounded-md
					p-5">One or more files already exist</span>`)
				return
			} else {
				http.Error(w, "Failed to save new file", http.StatusInternalServerError)
				return
			}
		}
		defer dst.Close()

		bytes, err := io.Copy(dst, file)
		fmt.Printf("%d bytes successfully copied to new file\n", bytes)

		uploadedFiles = append(uploadedFiles, fileHeader.Filename)
	}
	w.WriteHeader(http.StatusCreated) // Set the correct HTTP status code
	fmt.Fprintf(w, `<span remove-me="2s"
		class="bg-green-500
		text-white
		rounded-md
		p-5">New file(s) added</span>`)
}
