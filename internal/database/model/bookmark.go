package model

import (
	"time"

	"github.com/google/uuid"
)

type Visibility int
type TimelineKind string

const (
	VisibilityPublic Visibility = iota
	VisibilityPrivate
	VisibilityGroup
)

func (v Visibility) String() string {
	switch v {
	case VisibilityPublic:
		return "public"
	case VisibilityPrivate:
		return "private"
	case VisibilityGroup:
		return "group"
	}

	return ""
}

const (
	TimelineUser  TimelineKind = "user"
	TimelineGroup TimelineKind = "group"
)

type BookmarkRepository interface {
	Timeline(id uuid.UUID) ([]Bookmark, error)
	Get(id uuid.UUID) (Bookmark, error)
	Jump(id uuid.UUID) error
}

type Bookmark struct {
	ID          uuid.UUID  `json:"id"`
	OwnerID     uuid.UUID  `json:"-"`
	Owner       User       `json:"owner"`
	URL         string     `json:"url"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	CreatedAt   time.Time  `json:"created_at"`
}
