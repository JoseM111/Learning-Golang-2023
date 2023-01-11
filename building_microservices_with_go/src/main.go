package main

import (
	"context"
	"log"
	. "net/http"
	"os"
	"os/signal"
	"time"
	
	"main/handlers"
)

func main() {
	// logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	// handlers
	productHandler := handlers.NewProducts(l)
	
	// register handler as my service
	serveMux := NewServeMux()
	serveMux.Handle("/", productHandler)
	
	// creating our own server
	server := &Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	
	/*
		A `go` statement starts the execution of a function call as
		an independent concurrent thread of control, or `goroutine`,
		within the same address space.
	*/
	// won't block code because it is wrapped in a go function
	go func() {
		const portURL = "http://localhost:9090/"
		
		l.Printf("\nStarting server on port: %s", portURL)
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			// Exit causes the current program to exit with the given status code.
			// Conventionally, code zero indicates success, non-zero an error.
			// The program terminates immediately;
			// deferred functions are not run.
			os.Exit(1)
		}
	}()
	
	// channel
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	log.Println("Received terminate, graceful shutdown", sig)
	
	// gracefully shutdown
	timeOutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := server.Shutdown(timeOutContext)
	if err != nil {
		return
	}
}
