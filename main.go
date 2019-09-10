package main

import (
	"net/http"

	"github.com/kazu1029/imageUploader/uploader"
)

func setupRoutes() {
	http.HandleFunc("/upload", uploader.UploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRoutes()
}
