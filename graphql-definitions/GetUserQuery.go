package graphql_definitions

import (
	"github.com/chalkedgoose/act-up-api/entity"
	"github.com/graphql-go/graphql"
)

var GetUserQuery = &graphql.Field{
	Type:        UserType,
	Description: "Get a single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		return entity.User{
			ID:        id,
			Name:      "Carlos Alba",
			AvatarURL: "https://picsum.photos/350",
		}, nil
	},
}
