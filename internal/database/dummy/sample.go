package dummy

import (
	"time"

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
		Owner:       users[0],
		URL:         "https://twitch.tv/barveyhirdman",
		Description: "My Twitch channel",
		Visibility:  model.VisibilityPublic,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("49e76620-a86e-47a6-81d5-32cd9f7f8866"),
		Owner:       users[0],
		URL:         "https://twitter.com/barveyhirdman",
		Description: "My Twitter account",
		Visibility:  model.VisibilityPrivate,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a8"),
		Owner:       users[1],
		URL:         "https://google.com",
		Description: "Evil Search Engine",
		Visibility:  model.VisibilityPrivate,
		CreatedAt:   time.Now(),
	},
	{
		ID:          uuid.MustParse("914d44cc-d9fa-40ca-9d21-e90f705f45a9"),
		Owner:       users[1],
		URL:         "https://twitch.tv/rixraw",
		Description: "Good Twitch channel",
		Visibility:  model.VisibilityPublic,
		CreatedAt:   time.Now(),
	},
}
