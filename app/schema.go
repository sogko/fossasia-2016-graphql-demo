package app

import (
	"github.com/graphql-go/graphql"
)

/**

Let's try to write a blueprint for our GraphQL schema using the GraphQL Type System notation:

type User {
	id: Int
	name: String
	image_url: String
	nickname: String
	avatar_url: String
	age: Int
	is_single: Boolean
	best_friends: [User]
}

type Post {
	id: Int
	title: String
	author: User
	liked: [User]
	likedCount: Int
	viewCount: Int
}
type RootQuery {
	posts: [Post]
	users: [User]
}
*/

var UserType *graphql.Object
var PostType *graphql.Object
var Schema graphql.Schema

func init() {

	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.Int,
					Description: "User ID",
				},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "User's full name",
				},
				"nickname": &graphql.Field{
					Type:        graphql.String,
					Description: "User's nickname",
				},
				"avatar_url": &graphql.Field{
					Type:        graphql.String,
					Description: "Relative URL to user's avatar",
				},
				"age": &graphql.Field{
					Type:        graphql.Int,
					Description: "User's age",
				},
				"is_single": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Is user single?",
				},
				"best_friends": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "User's best friends",
				},
			}
		}),
	})

	PostType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Post",
		Description: "Post object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.Int,
					Description: "Post ID",
				},
				"title": &graphql.Field{
					Type:        graphql.String,
					Description: "Title of the post",
				},
				"author": &graphql.Field{
					Type:        UserType,
					Description: "Author of the post",
				},
				"liked": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "List of users who liked the post",
				},
				"viewCount": &graphql.Field{
					Type:        graphql.Int,
					Description: "Number of users who viewed the post",
				},
			}
		}),
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RootQuery",
			Description: "Root for all query objects on our GraphQL server",
			Fields: graphql.Fields{
				"posts": &graphql.Field{
					Type:        graphql.NewList(PostType),
					Description: "All the posts",
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `limit` arg
						limit := 10
						if v, ok := p.Args["limit"].(int); ok {
							limit = v
						}
						if limit > len(Posts) {
							limit = len(Posts)
						}
						return Posts[0:limit], nil
					},
				},
				"users": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "All the users",
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// get value of `limit` arg
						limit := 10
						if v, ok := p.Args["limit"].(int); ok {
							limit = v
						}
						if limit > len(Users) {
							limit = len(Users)
						}
						return Users[0:limit], nil
					},
				},
			},
		}),
	})
	if err != nil {
		panic(err)
	}
}
