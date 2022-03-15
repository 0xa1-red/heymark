package dummy

import (
	"context"
	"fmt"
	"sync"

	_context "github.com/alfreddobradi/heymark/internal/context"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/davecgh/go-spew/spew"
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

func (db *DummyDB) Timeline(ctx context.Context) ([]model.Bookmark, error) {
	db.Bookmarks.mx.Lock()
	defer db.Bookmarks.mx.Unlock()

	id := uuid.Nil
	authData, err := _context.GetAuthData(ctx)
	if err != nil && err.Error() != _context.ErrNoAuthorization.Error() {
		return nil, err
	} else if err == nil {
		user, err := db.Authorize(authData)
		if err != nil {
			return nil, err
		}

		id = user.ID
	}

	spew.Dump(id)

	res := []model.Bookmark{}
	for _, bookmark := range db.Bookmarks.Records {
		if bookmark.Owner.ID.String() == id.String() || bookmark.Visibility == model.VisibilityPublic {
			res = append(res, bookmark)
		}
	}

	return res, nil
}

func (db *DummyDB) GetBookmark(ctx context.Context, id uuid.UUID) (model.Bookmark, error) {
	db.Bookmarks.mx.Lock()
	defer db.Bookmarks.mx.Unlock()

	userID := uuid.Nil
	authData, err := _context.GetAuthData(ctx)
	if err != nil && err.Error() != _context.ErrNoAuthorization.Error() {
		return model.Bookmark{}, err
	} else if err == nil {
		user, err := db.authorize(authData)
		if err != nil {
			return model.Bookmark{}, err
		}

		userID = user.ID
	}

	bookmark, ok := db.Bookmarks.Records[id]
	if !ok {
		return model.Bookmark{}, fmt.Errorf("Bookmark not found")
	}

	if userID.String() == uuid.Nil.String() && bookmark.Visibility != model.VisibilityPublic {
		return model.Bookmark{}, model.ErrUnauthorized
	}

	return bookmark, nil
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
