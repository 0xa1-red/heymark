package schema

import (
	"fmt"

	"github.com/alfreddobradi/heymark/internal/database"
	"github.com/alfreddobradi/heymark/internal/database/model"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var TokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
	},
})

var userQueryFields = graphql.Fields{
	"user": &graphql.Field{
		Name:        "User",
		Type:        UserType,
		Description: "Return user details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal ID")
			}

			uid, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("Error parsing UUID: %w", err)
			}
			db := database.GetDB()
			user, err := db.GetUser(uid)
			if err != nil {
				return nil, fmt.Errorf("Error looking up user: %w", err)
			}
			return user, nil
		},
	},
	"authenticate": &graphql.Field{
		Name:        "Authenticate",
		Type:        TokenType,
		Description: "Exchange login details to a bearer token",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			username, ok := params.Args["username"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal username")
			}
			password, ok := params.Args["password"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal password")
			}

			db := database.GetDB()
			token, err := db.Authenticate(username, password)
			if err != nil {
				return nil, fmt.Errorf("Error looking up user: %w", err)
			}
			return token, nil
		},
	},
	"me": &graphql.Field{
		Name:        "Me",
		Type:        UserType,
		Description: "Return details for logged in user",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			authHeader := params.Context.Value(ContextAuth("Authorization"))
			if authHeader == nil {
				return nil, model.ErrUnauthorized
			}

			authData, err := model.GetAuthDataFromHeader(authHeader.(string))
			if err != nil {
				return nil, fmt.Errorf("Error parsing authorization data: %w", err)
			}

			spew.Dump(authData)

			db := database.GetDB()
			user, err := db.Authorize(authData)
			if err != nil {
				return nil, model.ErrUnauthorized
			}

			return user, nil
		},
	},
}

var userMutationFields = graphql.Fields{
	"user": &graphql.Field{
		Name:        "Create",
		Type:        UserType,
		Description: "Create a new user",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Username",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "Password",
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			username, ok := params.Args["username"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal username")
			}
			password, ok := params.Args["password"].(string)
			if !ok {
				return nil, fmt.Errorf("Failed to unmarshal password")
			}

			db := database.GetDB()
			user, err := db.CreateUser(username, password)
			if err != nil {
				return nil, fmt.Errorf("Failed to create user: %w", err)
			}

			return user, nil
		},
	},
}
