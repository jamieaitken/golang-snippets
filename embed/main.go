package main

import (
	"embed"
	"io/fs"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", loadEmbeddedContent())
}

//go:embed website/*
var website embed.FS

func loadEmbeddedContent () http.Handler {
	handler, err := fs.Sub(website, "website")
	if err != nil {
		return http.NotFoundHandler()
	}

	return http.FileServer(http.FS(handler))
}

