package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	
	"golangmovies.com/rating/internal/controller/rating"
	"golangmovies.com/rating/internal/repository"
	model "golangmovies.com/rating/pkg"
)

// ==================================================

// Handler defines a rating service controller.
type Handler struct {
	ctrl *rating.Controller
}

// ==================================================

// New creates a new rating service HTTP handler. (constructor)
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}

// ==================================================

// Handle handles PUT and GET /rating requests.
func (handler *Handler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	// FormValue returns the first value for the named component of the query.
	// POST and PUT body parameters take precedence over URL query string values.
	recordID := model.RecordID(request.FormValue("id"))
	if recordID == "" {
		// STATUS CODE: 400
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	
	recordType := model.RecordType(request.FormValue("type"))
	if recordType == "" {
		// STATUS CODE: 400
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	
	switch request.Method {
	case http.MethodGet:
		{
			v, err := handler.ctrl.GetAggregatedRating(request.Context(), recordID, recordType)
			if err != nil && errors.Is(err, repository.ErrNotFound) {
				responseWriter.WriteHeader(http.StatusNotFound)
				return
			}
			
			if err := json.NewEncoder(responseWriter).Encode(v); err != nil {
				log.Printf("Response encode error: %#v\n", err)
			}
		}
	case http.MethodPut:
		{
			userID := model.UserID(request.FormValue("userId"))
			value, err := strconv.ParseFloat(request.FormValue("value"), 64)
			
			if err != nil {
				// STATUS CODE: 400
				responseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			
			if err := handler.ctrl.PutRating(request.Context(), recordID, recordType, &model.Rating{
				UserID: userID,
				Value:  model.RatingValue(value),
			}); err != nil {
				log.Printf("Repository put error: %#v\n", err)
				// STATUS CODE: 500
				responseWriter.WriteHeader(http.StatusInternalServerError)
			}
		}
	default:
		// STATUS CODE: 400
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	
}

// ==================================================
// ==================================================
