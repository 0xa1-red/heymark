package dummy

import (
	"fmt"
	"sync"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/google/uuid"
)

type bookmarkRepository struct {
	mx      *sync.Mutex
	records map[uuid.UUID]model.Bookmark
}

func (db *DummyDB) Timeline(id uuid.UUID) ([]model.Bookmark, error) {
	return nil, fmt.Errorf("not implemented")
}

func (db *DummyDB) Get(id uuid.UUID) (model.Bookmark, error) {
	return model.Bookmark{}, fmt.Errorf("not implemented")
}

func (db *DummyDB) Jump(id uuid.UUID) error {
	return fmt.Errorf("not implemented")
}
