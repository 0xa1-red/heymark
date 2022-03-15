package model

import (
	"fmt"
	"net/url"
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
	VisibilityPublicLabel  string = "public"
	VisibilityPrivateLabel string = "private"
	VisibilityGroupLabel   string = "group"
)

func (v Visibility) String() string {
	switch v {
	case VisibilityPublic:
		return VisibilityPublicLabel
	case VisibilityPrivate:
		return VisibilityPrivateLabel
	case VisibilityGroup:
		return VisibilityGroupLabel
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
	CreateBookmark(owner User, bookmark Bookmark) (Bookmark, error)
}

type Bookmark struct {
	ID          uuid.UUID  `json:"id"`
	Owner       User       `json:"owner"`
	URL         string     `json:"url"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	CreatedAt   time.Time  `json:"created_at"`
}

func NewBookmark(owner User, bookmarkURL, description, visibility string) (Bookmark, error) {
	validatedURL, err := url.ParseRequestURI(bookmarkURL)
	if err != nil {
		return Bookmark{}, fmt.Errorf("Invalid URL (%s): %w", bookmarkURL, err)
	}

	var validatedVisibility Visibility
	switch visibility {
	case VisibilityPublicLabel:
		validatedVisibility = VisibilityPublic
	case VisibilityPrivateLabel:
		validatedVisibility = VisibilityPrivate
	case VisibilityGroupLabel:
		validatedVisibility = VisibilityGroup
	default:
		return Bookmark{}, fmt.Errorf("Invalid visibility: %s", visibility)
	}

	return Bookmark{
		ID:          uuid.New(),
		Owner:       owner,
		URL:         validatedURL.String(),
		Description: description,
		Visibility:  validatedVisibility,
		CreatedAt:   time.Now(),
	}, nil
}
