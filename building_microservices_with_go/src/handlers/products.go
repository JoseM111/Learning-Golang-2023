package handlers

import (
	"log"
	. "net/http"
	
	"main/data"
)

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (receiver *Products) ServeHTTP(responseWriter ResponseWriter, request *Request) {
	if request.Method == MethodGet {
		receiver.getProducts(responseWriter, request)
	}
	
	// catch-all STATUS CODE: 405-->Method Not Allowed
	responseWriter.WriteHeader(StatusMethodNotAllowed)
}

// getProducts GET REQUEST
func (receiver *Products) getProducts(responseWriter ResponseWriter, request *Request) {
	listOfProducts := data.GetProducts()
	
	// returning the list in a json format
	err := listOfProducts.ToJson(responseWriter)
	if err != nil {
		// STATUS CODE: 500
		Error(
			responseWriter, "Unable to marshal json",
			StatusInternalServerError,
		)
	}
}

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
