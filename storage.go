package main

import (
	"net/http"
)

func handleStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Served-With-Love-By", "Gophers")

	http.FileServer(http.Dir("./womenwhogo.org/")).ServeHTTP(w, r)
}

func handleAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=86400")
	handleStatic(w, r)
}
