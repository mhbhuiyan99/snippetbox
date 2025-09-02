package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)


func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// structured logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	mux := http.NewServeMux();

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on '/' only
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}