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
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// handlers
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	
	// register handler as my service
	serveMux := NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	
	// creating our own server
	server := &Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	/*
		A `go` statement starts the execution of a function call as
		an independent concurrent thread of control, or `goroutine`,
		within the same address space.
	*/
	// won't block code because it is wrapped in a go function
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
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
	_ = server.Shutdown(timeOutContext)
}
