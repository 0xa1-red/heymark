package dummy

import (
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/google/uuid"
)

func (db *DummyDB) Timeline(id uuid.UUID) ([]model.Bookmark, error) {
	return nil, fmt.Errorf("not implemented")
}

func (db *DummyDB) Get(id uuid.UUID) (model.Bookmark, error) {
	return model.Bookmark{}, fmt.Errorf("not implemented")
}

func (db *DummyDB) Jump(id uuid.UUID) error {
	return fmt.Errorf("not implemented")
}
