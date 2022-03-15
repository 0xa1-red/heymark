package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

const (
	FieldQuery    string = "query"
	FieldMutation string = "mutation"
)

var queryFields graphql.Fields = graphql.Fields{}
var mutationFields graphql.Fields = graphql.Fields{}

func registerFields(fields graphql.Fields, kind string, overwrite ...bool) {
	ow := false
	if len(overwrite) > 0 {
		ow = overwrite[0]
	}

	var collection graphql.Fields
	switch kind {
	case FieldQuery:
		collection = queryFields
	case FieldMutation:
		collection = mutationFields
	default:
		return
	}

	for key, field := range fields {
		if _, ok := collection[key]; ok {
			if ow {
				log.Printf("Query field %s already exists. Overwriting...", key)
				collection[key] = field
			} else {
				log.Printf("Query field %s already exists. Skipping...", key)
			}
			continue
		}

		collection[key] = field
	}
}

var RootSchema graphql.Schema

func init() {
	registerFields(userQueryFields, FieldQuery)
	registerFields(userMutationFields, FieldMutation)
	registerFields(bookmarkQueryFields, FieldQuery)
	registerFields(bookmarkMutationFields, FieldMutation)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Operations to read data",
		Fields:      queryFields,
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Operations to change data",
		Fields:      mutationFields,
	})

	RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
