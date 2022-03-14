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

	userCache := map[uuid.UUID]model.User{}

	res := []model.Bookmark{}
	for _, bookmark := range db.Bookmarks.Records {
		if bookmark.OwnerID == id || bookmark.Visibility == model.VisibilityPublic {
			var user model.User
			var cacheHit bool
			user, cacheHit = userCache[bookmark.OwnerID]
			if !cacheHit {
				var ok bool
				user, ok = db.Users.Records[bookmark.OwnerID]
				if !ok {
					return nil, fmt.Errorf("User with ID %s not found", bookmark.OwnerID)
				}
				userCache[bookmark.OwnerID] = user
			}
			bookmark.Owner = user
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
