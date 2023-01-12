package handlers

import (
	. "fmt"
	"io/ioutil"
	"log"
	. "net/http"
)

// ##############################################

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log: log}
}

// #################### HANDLERS #######################

func (h *Hello) ServeHTTP(writer ResponseWriter, request *Request) {
	h.log.Println("Hola MUndo•..")
	
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		Error(writer, "Oops! Something went wrong", StatusBadRequest)
	}
	
	_, _ = Fprintf(writer, "Hello %s\n", data)
}

// ##############################################
