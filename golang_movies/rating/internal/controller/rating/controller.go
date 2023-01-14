package rating

import (
	"context"
	"errors"
	
	"golangmovies.com/rating/internal/repository"
	model "golangmovies.com/rating/pkg"
)

// ================================================

// ratingRepository interface that is private to this file
type ratingRepository interface {
	Get(
		ctx context.Context,
		recordID model.RecordID,
		recordType model.RecordType) ([]model.Rating, error)
	
	Put(
		ctx context.Context,
		recordID model.RecordID,
		recordType model.RecordType,
		rating *model.Rating) error
}

// ================================================

// Controller defines a rating service controller.
type Controller struct {
	repo ratingRepository
}

// ================================================

// New creates a rating service controller. (constructor)
func New(repo ratingRepository) *Controller {
	return &Controller{repo}
}

// ================================================

// GetAggregatedRating returns the aggregated rating for a
// record or ErrNotFound if there are no ratings for it.
func (ctrl *Controller) GetAggregatedRating(
	ctx context.Context,
	recordID model.RecordID,
	recordType model.RecordType) (float64, error) {
	
	ratings, err := ctrl.repo.Get(ctx, recordID, recordType)
	if err != nil && err == repository.ErrNotFound {
		return 0, errors.New("ratings not found for a record")
	} else if err != nil {
		return 0, err
	}
	
	sum := float64(0)
	
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	
	return sum / float64(len(ratings)), nil
}

// PutRating writes a rating for a given record.
func (ctrl *Controller) PutRating(
	ctx context.Context,
	recordID model.RecordID,
	recordType model.RecordType,
	rating *model.Rating) error {
	
	return ctrl.repo.Put(ctx, recordID, recordType, rating)
}

// ================================================
