package dummy_test

import (
	"testing"
	"time"

	"github.com/alfreddobradi/heymark/internal/database/dummy"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/alfreddobradi/heymark/internal/helper"
	"github.com/google/uuid"
)

var users = []model.User{
	{
		ID:       uuid.MustParse("f8abde0c-7f69-4978-8beb-d5253252e2a1"),
		Username: "Barvey",
		Password: helper.Sha256("test123"),
		Bio:      "Hello I am gamer",
	},
	{
		ID:       uuid.MustParse("b146d2cb-94c7-45be-bcda-f047aa3dee04"),
		Username: "FellowGamer",
		Password: helper.Sha256("test1234"),
		Bio:      "Hello I am too a gamer",
	},
}

var bookmarks = []model.Bookmark{
	{
		ID:          uuid.MustParse("10854e1b-53d7-4d85-aa43-13bde0601729"),
		OwnerID:     users[0].ID,
		URL:         "https://twitch.tv/barveyhirdman",
		Description: "My Twitch channel",
		Visibility:  model.VisibilityPublic,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("49e76620-a86e-47a6-81d5-32cd9f7f8866"),
		OwnerID:     users[0].ID,
		URL:         "https://twitter.com/barveyhirdman",
		Description: "My Twitter account",
		Visibility:  model.VisibilityPrivate,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a8"),
		OwnerID:     users[1].ID,
		URL:         "https://google.com",
		Description: "Evil Search Engine",
		Visibility:  model.VisibilityPrivate,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a9"),
		OwnerID:     users[1].ID,
		URL:         "https://twitch.tv/rixraw",
		Description: "Good Twitch channel",
		Visibility:  model.VisibilityPublic,
		CreatedAt:   time.Now(),
	},
}

func CreateTestDB() *dummy.DummyDB {

	userRepository := dummy.NewUserRepository()
	for _, user := range users {
		userRepository.Records[user.ID] = user
	}

	bookmarkRepository := dummy.NewBookmarkRepository()
	for _, bookmark := range bookmarks {
		bookmarkRepository.Records[bookmark.ID] = bookmark
	}

	return &dummy.DummyDB{
		Users:     userRepository,
		Bookmarks: bookmarkRepository,
	}
}

func assertIDInTimeline(id uuid.UUID, timeline []model.Bookmark) bool {
	for _, bookmark := range timeline {
		if id.String() == bookmark.ID.String() {
			return true
		}
	}
	return false
}

func TestTimeline(t *testing.T) {
	sampleDB := CreateTestDB()

	tests := []struct {
		label       string
		id          uuid.UUID
		expectedIDs []uuid.UUID
	}{
		{
			label: "Public timeline",
			id:    uuid.Nil,
			expectedIDs: []uuid.UUID{
				bookmarks[0].ID, bookmarks[3].ID,
			},
		},
		{
			label: "Timeline for first user",
			id:    users[0].ID,
			expectedIDs: []uuid.UUID{
				bookmarks[0].ID, bookmarks[1].ID, bookmarks[3].ID,
			},
		},
		{
			label: "Timeline for second user",
			id:    users[1].ID,
			expectedIDs: []uuid.UUID{
				bookmarks[0].ID, bookmarks[2].ID, bookmarks[3].ID,
			},
		},
	}

	for i, tt := range tests {
		tf := func(t *testing.T) {
			t.Logf("%d - %s", i+1, tt.label)

			timeline, err := sampleDB.Timeline(tt.id)
			if err != nil {
				t.Errorf("Error: couldn't get timeline: %v", err)
			}

			for _, expected := range tt.expectedIDs {
				if !assertIDInTimeline(expected, timeline) {
					t.Fatalf("Fail: expected ID %s to be in timeline, but it isn't", expected.String())
				}
			}

			t.Log("Pass")
		}

		t.Run(tt.label, tf)
	}
}
