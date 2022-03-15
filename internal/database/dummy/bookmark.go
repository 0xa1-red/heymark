package dummy

import (
	"fmt"
	"sync"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/google/uuid"
)

type BookmarkRepository struct {
	mx      *sync.Mutex
	Records map[uuid.UUID]model.Bookmark
}

func NewBookmarkRepository() BookmarkRepository {
	return BookmarkRepository{
		mx:      &sync.Mutex{},
		Records: map[uuid.UUID]model.Bookmark{},
	}
}

func (db *DummyDB) Timeline(id uuid.UUID) ([]model.Bookmark, error) {
	db.Bookmarks.mx.Lock()
	db.Users.mx.Lock()
	defer db.Bookmarks.mx.Unlock()
	defer db.Users.mx.Unlock()

	res := []model.Bookmark{}
	for _, bookmark := range db.Bookmarks.Records {
		if bookmark.Owner.ID.String() == id.String() || bookmark.Visibility == model.VisibilityPublic {
			res = append(res, bookmark)
		}
	}

	return res, nil
}

func (db *DummyDB) Get(id uuid.UUID) (model.Bookmark, error) {
	return model.Bookmark{}, fmt.Errorf("not implemented")
}

func (db *DummyDB) Jump(id uuid.UUID) error {
	return fmt.Errorf("not implemented")
}

func (db *DummyDB) CreateBookmark(owner model.User, bookmark model.Bookmark) (model.Bookmark, error) {
	db.Bookmarks.mx.Lock()
	defer db.Bookmarks.mx.Unlock()

	for _, existing := range db.Bookmarks.Records {
		if existing.URL == bookmark.URL && existing.Owner.ID == owner.ID {
			return model.Bookmark{}, fmt.Errorf("The URL is already bookmarked: %s", existing.URL)
		}
	}

	db.Bookmarks.Records[bookmark.ID] = bookmark

	return bookmark, nil
}
