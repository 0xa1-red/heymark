package dummy

import (
	"sync"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
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
	uids := []uuid.UUID{
		uuid.MustParse("f8abde0c-7f69-4978-8beb-d5253252e2a1"),
		uuid.MustParse("b146d2cb-94c7-45be-bcda-f047aa3dee04"),
	}
	users := UserRepository{
		mx: &sync.Mutex{},
		Records: map[uuid.UUID]model.User{
			uids[0]: {
				ID:       uids[0],
				Username: "Barvey",
				Password: helper.Sha256("test123"),
				Bio:      "Hello I am gamer",
			},
			uids[1]: {
				ID:       uids[1],
				Username: "FellowGamer",
				Password: helper.Sha256("test1234"),
				Bio:      "Hello I am too a gamer",
			},
		},
	}

	bids := []uuid.UUID{
		uuid.MustParse("10854e1b-53d7-4d85-aa43-13bde0601729"),
		uuid.MustParse("49e76620-a86e-47a6-81d5-32cd9f7f8866"),
		uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a8"),
		uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a9"),
	}
	bookmarks := BookmarkRepository{
		mx: &sync.Mutex{},
		Records: map[uuid.UUID]model.Bookmark{
			bids[0]: {
				ID:          bids[0],
				OwnerID:     uids[0],
				URL:         "https://twitch.tv/barveyhirdman",
				Description: "My Twitch channel",
				Visibility:  model.VisibilityPublic,
				CreatedAt:   time.Now(),
			},
			bids[1]: {
				ID:          bids[1],
				OwnerID:     uids[0],
				URL:         "https://twitter.com/barveyhirdman",
				Description: "My Twitter account",
				Visibility:  model.VisibilityPrivate,
				CreatedAt:   time.Now(),
			},
			bids[2]: {
				ID:          bids[2],
				OwnerID:     uids[1],
				URL:         "https://google.com",
				Description: "Evil Search Engine",
				Visibility:  model.VisibilityPrivate,
				CreatedAt:   time.Now(),
			},
			bids[3]: {
				ID:          bids[3],
				OwnerID:     uids[1],
				URL:         "https://twitch.tv/rixraw",
				Description: "Good Twitch channel",
				Visibility:  model.VisibilityPublic,
				CreatedAt:   time.Now(),
			},
		},
	}

	db.Users = users
	db.Bookmarks = bookmarks
}
