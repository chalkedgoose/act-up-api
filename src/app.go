package main

import (
	"github.com/chalkedgoose/act-up-api/src/handler"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

var fields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}

var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

func main() {
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}
