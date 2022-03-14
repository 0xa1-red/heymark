package schema

import (
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/graphql-go/graphql"
)

var bookmarkQueryFields = graphql.Fields{
	"timeline": &graphql.Field{
		Type: BookmarkType,
		Args: graphql.FieldConfigArgument{
			// TODO filters and sorting
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			authHeader := params.Context.Value(ContextAuth("Authorization"))
			if authHeader == nil {
				return nil, model.ErrUnauthorized
			}

			authData, err := model.GetAuthDataFromHeader(authHeader.(string))
			if err != nil {
				return nil, fmt.Errorf("Error parsing authorization data: %w", err)
			}

			db := database.GetDB()
			user, err := db.Authorize(authData)
			if err != nil {
				return nil, model.ErrUnauthorized
			}

			timeline, err := db.Timeline(user.ID)
			if err != nil {
				return nil, fmt.Errorf("Error retrieving timeline: %w", err)
			}
			return timeline, nil
		},
	},
}
