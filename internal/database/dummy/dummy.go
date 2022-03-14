package dummy

import (
	"sync"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
)

type DummyDB struct {
	Users     userRepository
	Bookmarks bookmarkRepository
}

var _ model.UserRepository = &DummyDB{}
var _ model.BookmarkRepository = &DummyDB{}

func New() *DummyDB {
	uid := uuid.MustParse("f8abde0c-7f69-4978-8beb-d5253252e2a1")
	users := userRepository{
		mx: &sync.Mutex{},
		records: map[uuid.UUID]model.User{
			uid: {
				ID:       uid,
				Username: "Barvey",
				Password: helper.Sha256("test123"),
				Bio:      "Hello I am gamer",
			},
		},
	}

	bid := uuid.MustParse("10854e1b-53d7-4d85-aa43-13bde0601729")
	bookmarks := bookmarkRepository{
		mx: &sync.Mutex{},
		records: map[uuid.UUID]model.Bookmark{
			bid: {
				ID:          bid,
				OwnerID:     uid,
				URL:         "https://twitch.tv/barveyhirdman",
				Description: "My Twitch channel",
				Visibility:  model.VisibilityPublic,
				CreatedAt:   time.Now(),
			},
		},
	}

	return &DummyDB{
		Users:     users,
		Bookmarks: bookmarks,
	}
}
