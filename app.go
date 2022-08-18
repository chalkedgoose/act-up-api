package main

import (
	"github.com/chalkedgoose/act-up-api/graphql-definitions"
	"github.com/chalkedgoose/act-up-api/handler"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

func main() {
	newSchema, err := graphql.NewSchema(graphql_definitions.AppSchemaConfig)

	if err != nil {
		log.Fatalf("failed to create new newSchema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &newSchema,
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
