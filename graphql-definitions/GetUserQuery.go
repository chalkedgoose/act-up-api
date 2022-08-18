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

var GetUsersQuery = &graphql.Field{
	Type:        graphql.NewList(UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		users := []entity.User{
			entity.User{
				ID:        "1",
				Name:      "Carlos Alba",
				AvatarURL: "https://picsum.photos/350",
			},
			entity.User{
				ID:        "2",
				Name:      "Haley Levesque",
				AvatarURL: "https://picsum.photos/350",
			},
			entity.User{
				ID:        "3",
				Name:      "Kit Alba",
				AvatarURL: "https://picsum.photos/350",
			},
			entity.User{
				ID:        "4",
				Name:      "Pablo Alba",
				AvatarURL: "https://picsum.photos/350",
			},
		}

		return users, nil
	},
}
