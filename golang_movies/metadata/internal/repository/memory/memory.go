package memory

import (
	. "context"
	"sync"
	
	"golangmovies.com/metadata/internal/repository"
	model "golangmovies.com/metadata/pkg"
)

// =============================================

// Repository defines a memory movie metadata repository.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// =============================================

// New creates a new memory repository. (constructor)
func New() *Repository {
	return &Repository{
		// RWMutex implementation is using a sync.RWMutex
		// structure to protect against concurrent writes and reads.
		RWMutex: sync.RWMutex{},
	}
}

// ====================== CRUD-FUNCTIONS =======================

// Get retrieve's movie metadata for by movie.
func (repo *Repository) Get(_ Context, id string) (*model.Metadata, error) {
	repo.RLock()
	// undo a single RLock call
	defer repo.RUnlock()
	
	metadata, ok := repo.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	
	return metadata, nil
}

// Put adds movie metadata for a given movie id
func (repo *Repository) Put(_ Context, id string, metadata *model.Metadata) error {
	repo.RLock()
	// undo a single RLock call
	defer repo.RUnlock()
	
	repo.data[id] = metadata
	return nil
}

// =============================================
