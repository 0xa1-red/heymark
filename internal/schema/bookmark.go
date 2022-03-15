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
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"owner": &graphql.Field{
			Type: UserType,
		},
		"visibility": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var bookmarkQueryFields = graphql.Fields{
	"timeline": &graphql.Field{
		Type: graphql.NewList(BookmarkType),
		Args: graphql.FieldConfigArgument{
			// TODO filters and sorting
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			authHeader := params.Context.Value(ContextAuth("Authorization"))

			header, _ := authHeader.(string)
			authData, err := model.GetAuthDataFromHeader(header)
			if err != nil {
				return nil, fmt.Errorf("Error parsing authorization data: %w", err)
			}

			db := database.GetDB()
			id := uuid.Nil
			if authData.Username != "" {
				user, err := db.Authorize(authData)
				if err != nil {
					return nil, model.ErrUnauthorized
				}
				id = user.ID
			}
			timeline, err := db.Timeline(id)
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
				Description: "Visibility (public|private)",
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			authHeader := params.Context.Value(ContextAuth("Authorization"))
			if authHeader == nil {
				return nil, model.ErrUnauthorized
			}

			header, ok := authHeader.(string)
			if !ok {
				return nil, model.ErrUnauthorized
			}

			authData, err := model.GetAuthDataFromHeader(header)
			if err != nil {
				return nil, fmt.Errorf("Error parsing authorization data: %w", err)
			}

			db := database.GetDB()
			user, err := db.Authorize(authData)
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
