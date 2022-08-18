package main

import (
	"github.com/chalkedgoose/act-up-api/src/handler"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.Any("/graphql", gin.WrapH(h))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run()

	if err != nil {
		return
	}
}
