package handlers

import (
	"log"
	. "net/http"
	"regexp"
	"strconv"

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
	switch verb := request.Method; verb {
	case MethodGet:
		receiver.getProducts(responseWriter, request)
	case MethodPost:
		receiver.addProducts(responseWriter, request)
	case MethodPut:
		{
			// TODO: handle PUT requests
			// expect the ID in the URI.
			// NOTE the other frameworks handle this type of stuff
			regex := regexp.MustCompile(`/([0-9]+)`)
			oneGroup := regex.FindAllStringSubmatch(request.URL.Path, -1)

			if len(oneGroup) != 1 {
				receiver.logger.Println("Invalid URI.. more than one id")
				// STATUS CODE: 400
				Error(responseWriter, "Invalid URI", StatusBadRequest)
				return
			}

			if len(oneGroup[0]) != 2 {
				receiver.logger.Println("Invalid URI.. more than one capture group")
				// STATUS CODE: 400
				Error(responseWriter, "Invalid URI", StatusBadRequest)
				return
			}

			captureIDString := oneGroup[0][1]
			// convert id into an integer
			id, err := strconv.Atoi(captureIDString)
			if err != nil {
				receiver.logger.Println(
					"Invalid URI.. unable to convert to number",
					captureIDString,
				)

				// STATUS CODE: 400
				Error(responseWriter, "Invalid URI", StatusBadRequest)
				return
			}
			// calling the put function
			receiver.updateProducts(id, responseWriter, request)
			receiver.logger.Println("got id:", id)
		}
	default:
		// catch-all STATUS CODE: 405-->Method Not Allowed
		responseWriter.WriteHeader(StatusMethodNotAllowed)
	}
}

// getProducts GET REQUEST
func (receiver *Products) getProducts(responseWriter ResponseWriter, request *Request) {
	receiver.logger.Println("Handle GET Products")

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

// addProducts POST REQUEST
func (receiver *Products) addProducts(responseWriter ResponseWriter, request *Request) {
	receiver.logger.Println("Handle POST Products")

	newProduct := &data.Product{}

	err := newProduct.FromJson(request.Body)
	if err != nil {
		// STATUS CODE: 400
		Error(
			responseWriter, "Unable to marshal json",
			StatusBadRequest,
		)
	}

	receiver.logger.Printf("\nProduct: %#v", newProduct)
	data.AddProducts(newProduct)
}

// updateProducts PUT REQUEST updates the data
func (receiver *Products) updateProducts(id int, responseWriter ResponseWriter, request *Request) {
	receiver.logger.Println("Handle PUT Products")

	updatedProduct := &data.Product{}

	err := updatedProduct.FromJson(request.Body)
	if err != nil {
		// STATUS CODE: 400
		Error(
			responseWriter, "Unable to marshal json",
			StatusBadRequest,
		)
	}

	err = data.UpdateProducts(id, updatedProduct)
	if err == data.ErrProductNotFound {
		// STATUS CODE: 404
		Error(responseWriter, "Product not found", StatusNotFound)
		return
	}

	if err != nil {
		// STATUS CODE: 500
		Error(responseWriter, "Product not found", StatusInternalServerError)
		return
	}
}

// ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
