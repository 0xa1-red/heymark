package schema

import (
	"github.com/graphql-go/graphql"
)

type ContextAuth string

// TODO implement groups
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Group",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Group ID (UUID)",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The group's name",
		},
		"url": &graphql.Field{
			Type:        graphql.String,
			Description: "The group's URL slug",
		},
		"created_at": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "The timestamp the group was created at",
		},
	},
})

// TODO implement roles and RBA
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Role ID (UUID)",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The role's name",
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: "A short description of the role",
		},
	},
})
