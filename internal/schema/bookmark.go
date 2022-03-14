package schema

import (
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/davecgh/go-spew/spew"
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

			spew.Dump(timeline)

			return timeline, nil
		},
	},
}
