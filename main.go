package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrerrh/local-drop/handlers"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func loadConfig(configFile string) (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	return config, err
}

func main() {
	os.MkdirAll("files", os.ModePerm) //Makes sure the "files" folder exists to save them
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download/", handlers.DownloadHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/files/", handlers.FilesListingHandler)

	//Image loading handler
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	config, err := loadConfig("config.json")
	if err != nil {
		log.Println("No config.json found, using default port 8080")
		config.Port = "8080"
	}

	http.Handle("/", http.FileServer(http.Dir("templates")))
	fmt.Println("Server running at port:" + config.Port)
	err = http.ListenAndServe("0.0.0.0:"+config.Port, nil)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", config.Port, err)
	}

}
