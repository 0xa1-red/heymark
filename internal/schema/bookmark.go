package schema

import (
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var BookmarkType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bookmark",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Bookmark ID (UUID)",
		},
		"url": &graphql.Field{
			Type:        graphql.String,
			Description: "Bookmark URL",
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: "A short description of the bookmark",
		},
		"owner": &graphql.Field{
			Type:        UserType,
			Description: "Details of the user who created the bookmark",
		},
		"visibility": &graphql.Field{
			Type:        graphql.String,
			Description: "Visibility of the bookmark (public, private or group)",
		},
		"created_at": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "The timestamp the bookmark was created at",
		},
	},
})

var bookmarkQueryFields = graphql.Fields{
	"bookmark": &graphql.Field{
		Type:        BookmarkType,
		Description: "Return a single bookmark",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The ID of the bookmark",
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal ID")
			}

			bid, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse bookmark ID: %w", err)
			}

			db := database.GetDB()
			bookmark, err := db.GetBookmark(params.Context, bid)
			if err != nil {
				return nil, err
			}

			return bookmark, nil
		},
	},
	"timeline": &graphql.Field{
		Type:        graphql.NewList(BookmarkType),
		Description: "Return a timeline of bookmarks",
		Args:        graphql.FieldConfigArgument{
			// TODO filters and sorting
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			db := database.GetDB()

			timeline, err := db.Timeline(params.Context)
			if err != nil {
				return nil, fmt.Errorf("Error retrieving timeline: %w", err)
			}

			return timeline, nil
		},
	},
}

var bookmarkMutationFields = graphql.Fields{
	"bookmark": &graphql.Field{
		Type: BookmarkType,
		Args: graphql.FieldConfigArgument{
			"url": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "URL you want to save",
			},
			"description": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "A short description",
			},
			"visibility": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Visibility (public or private)",
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			db := database.GetDB()
			user, err := authorize(params.Context, db)
			if err != nil {
				return nil, model.ErrUnauthorized
			}

			url, ok := params.Args["url"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal url")
			}
			description, ok := params.Args["description"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal description")
			}
			visibility, ok := params.Args["visibility"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal visibility")
			}

			bookmark, err := model.NewBookmark(user, url, description, visibility)
			if err != nil {
				return nil, fmt.Errorf("Error creating new bookmark: %w", err)
			}

			bookmark, err = db.CreateBookmark(user, bookmark)
			if err != nil {
				return nil, fmt.Errorf("Error saving bookmark timeline: %w", err)
			}

			return bookmark, nil
		},
	},
}
