package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type File struct {
	Id    int
	Name  string
	Size  float64
	SizeT string
}

func FilesListingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	filesPath, err := filepath.Abs("files/")
	if err != nil {
		http.Error(w, "Error getting the path to such file", http.StatusBadRequest)
		return
	}

	filesEnumPaths, err := filepath.Glob(filesPath + "/*")
	if err != nil {
		http.Error(w, "Error listing files", http.StatusBadRequest)
		return
	}

	var files []File

	currId := 0
	for _, f := range filesEnumPaths {
		file, err := os.Open(f)
		if err != nil {
			http.Error(w, "Error retrieving informations about file", http.StatusInternalServerError)
			fmt.Println("Error retrieving info of file: " + f)
			return
		}

		finfo, _ := file.Stat()
		defer file.Close()
		//Decide file units and truncate bytes value
		fileSize := finfo.Size()
		fileT := "Bytes"
		if (fileSize / (1 << 30)) >= 1 {
			fileT = "GBytes"
			fileSize /= (1 << 30)
		} else if (fileSize / (1 << 20)) >= 1 {
			fileT = "MBytes"
			fileSize /= (1 << 20)
		} else if (fileSize / (1 << 10)) >= 1 {
			fileT = "KBytes"
			fileSize /= (1 << 10)
		}

		files = append(files,
			File{
				Id:    currId,
				Name:  finfo.Name(),
				Size:  float64(fileSize),
				SizeT: fileT,
			})
		currId++
	}

	tmpl, err := template.ParseFiles("templates/files.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, files)

}
