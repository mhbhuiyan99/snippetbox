package main

import (
	"log"
	"net/http"
)


func main(){
	mux := http.NewServeMux();
	mux.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on '/' only
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}