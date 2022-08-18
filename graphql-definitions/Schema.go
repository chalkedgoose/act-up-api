package graphql_definitions

import (
	"github.com/graphql-go/graphql"
)

var fields = graphql.Fields{
	"user":  GetUserQuery,
	"users": GetUsersQuery,
}

var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

var AppSchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
