package memory

import (
	"context"
	
	"golangmovies.com/rating/internal/repository"
	model "golangmovies.com/rating/pkg"
)

// ================================================

// Repository defines a rating repository.
type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

// ================================================

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[model.RecordType]map[model.RecordID][]model.Rating{},
	}
}

// ====================== CRUD-FUNCTIONS =======================

// Get retrieves all ratings for a given record
func (repo *Repository) Get(
	ctx context.Context,
	recordID model.RecordID,
	recordType model.RecordType) ([]model.Rating, error) {
	
	if _, ok := repo.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	
	return repo.data[recordType][recordID], nil
}

// Put adds a rating for a given record.
func (repo *Repository) Put(
	_ context.Context, /* may change */
	recordID model.RecordID,
	recordType model.RecordType,
	rating *model.Rating) error {
	
	if _, ok := repo.data[recordType]; !ok {
		repo.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	
	dataList := repo.data[recordType][recordID]
	// The append built-in function appends elements to the end of a slice.
	// slice = append(slice, anotherSlice...)
	dataList = append(dataList, *rating)
	
	return nil
}

// ================================================
