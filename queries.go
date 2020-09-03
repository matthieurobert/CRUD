package main

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single user by id
			   http://localhost:8000/user?query={user(id:1){username,mail,password}}
			*/
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(u graphql.ResolveParams) (interface{}, error) {
					id, ok := u.Args["id"].(int)
					if ok {
						// find user
						for _, user := range users {
							if int(user.ID) == id {
								return user, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) user list
			   http://localhost:8000/user?query={list{id,username,mail,password}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get User List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
		},
	},
)
