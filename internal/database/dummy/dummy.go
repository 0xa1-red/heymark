package dummy

import (
	"github.com/alfreddobradi/heymark/internal/database/model"
)

type DummyDB struct {
	Users     UserRepository
	Bookmarks BookmarkRepository
}

var _ model.UserRepository = &DummyDB{}
var _ model.BookmarkRepository = &DummyDB{}

func New() *DummyDB {
	db := &DummyDB{
		Users:     UserRepository{},
		Bookmarks: BookmarkRepository{},
	}
	db.sampleData()
	return db
}

func (db *DummyDB) sampleData() {
	userRepository := NewUserRepository()
	for _, user := range users {
		userRepository.Records[user.ID] = user
	}

	bookmarkRepository := NewBookmarkRepository()
	for _, bookmark := range bookmarks {
		bookmarkRepository.Records[bookmark.ID] = bookmark
	}

	db.Users = userRepository
	db.Bookmarks = bookmarkRepository
}
