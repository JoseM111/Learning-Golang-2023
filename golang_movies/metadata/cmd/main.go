package main

import (
	"log"
	"net/http"
	
	"golangmovies.com/metadata/internal/controller/metadata"
	httphandler "golangmovies.com/metadata/internal/handler/http"
	"golangmovies.com/metadata/internal/repository/memory"
)

// main The function we just created initializes all structures of
// our service and starts the http API handler we implemented earlier.
// The service is ready to process user requests
func main() {
	log.Println("Starting the movie metadata service")
	
	repo := memory.New()
	ctrl := metadata.New(repo)
	
	handler := httphandler.New(ctrl)
	// The HandlerFunc type is an adapter to allow the use of ordinary
	// functions as HTTP handlers. If f is a function with the appropriate
	// signature, HandlerFunc(f) is a Handler that calls f.
	http.Handle("/metadata", http.HandlerFunc(handler.GetMetadata))
}
