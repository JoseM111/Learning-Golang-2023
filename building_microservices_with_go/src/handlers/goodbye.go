package handlers

import (
	"log"
	. "net/http"
)

// ###########################################

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log: log}
}

// #################### HANDLERS #######################
func (g *Goodbye) ServeHTTP(writer ResponseWriter, request *Request) {
	_, _ = writer.Write([]byte("Byee!!"))
}

// ###########################################
