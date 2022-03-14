package database

import (
	"github.com/alfreddobradi/heymark/internal/database/dummy"
	"github.com/alfreddobradi/heymark/internal/database/model"
)

type Kind string

const (
	KindDummy Kind = "dummy"
)

type Database interface {
	model.UserRepository
	model.BookmarkRepository
}

var db Database

func GetDB() Database {
	if db == nil {
		switch GetKind() {
		case KindDummy:
			db = dummy.New()
		}
	}
	return db
}
