package metadata

import (
	"context"
	"errors"
	
	"golangmovies.com/metadata/internal/repository"
	model "golangmovies.com/metadata/pkg"
)

// =============================================

// metadataRepository interface that is private to this file
type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// =============================================

// Controller defines a metadata service controller.
type Controller struct {
	repo metadataRepository
}

// =============================================

// New creates a metadata service controller. (constructor)
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

// ======================= MEMBER-FUNCTIONS ======================

// Get returns movie metadata by id.
func (repoController *Controller) Get(
	ctx context.Context, id string) (*model.Metadata, error) {
	
	response, err := repoController.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, errors.New("not found")
	}
	
	return response, err
}

// =============================================
