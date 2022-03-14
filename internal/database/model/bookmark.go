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
	OwnerID     uuid.UUID  `json:"owner_id"`
	URL         string     `json:"url"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	CreatedAt   time.Time  `json:"created_at"`
}
