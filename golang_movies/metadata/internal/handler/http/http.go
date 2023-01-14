package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	
	"golangmovies.com/metadata/internal/controller/metadata"
	"golangmovies.com/metadata/internal/repository"
)

// ================================================

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	ctrl *metadata.Controller
}

// ================================================

// New creates a new movie metadata HTTP handler. (constructor)
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl}
}

// ======================= MEMBER-FUNCTIONS =========================

// GetMetadata handles GET /metadata requests. The handler we just
// created uses our repository to retrieve the information and return
// it in JSON format. We chose JSON here just for simplicity.
func (handler *Handler) GetMetadata(responseWriter http.ResponseWriter, request *http.Request) {
	// FormValue returns the first value for the named component of the query
	id := request.FormValue("id")
	if id == "" {
		// STATUS CODE: 400
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// Context returns the request's context.
	ctx := request.Context()
	
	data, err := handler.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		// STATUS CODE: 404
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %#v\n", err)
		// STATUS CODE: 500
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	if err := json.NewEncoder(responseWriter).Encode(data); err != nil {
		log.Printf("Repository get error: %#v\n", err)
	}
}

// ================================================
