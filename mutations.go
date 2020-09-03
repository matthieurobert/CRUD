package main

import (
	"github.com/graphql-go/graphql"
	"math/rand"
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				/* Create new user item
				http://localhost:8000/user?query=mutation+_{create(username:"admin",mail:"admin@example.com",password:"root"){id,username,mail,password}}
				*/
				Type:        userType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"mail": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					user := User{
						ID:       int64(rand.Intn(100000)),
						Username: params.Args["username"].(string),
						Mail:     params.Args["mail"].(string),
						Password: params.Args["password"].(string),
					}
					users = append(users, user)
					return user, nil
				},
			},

			/* Update user by id
			http://localhost:8000/user?query=mutation+_{update(id:1,password:"root"){id,username,mail,password}}
			*/
			"update": &graphql.Field{
				Type:        userType,
				Description: "Update user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"mail": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					username, usernameOK := params.Args["username"].(string)
					mail, mailOK := params.Args["mail"].(string)
					password, passwordOK := params.Args["password"].(string)
					user := User{}
					for i, p := range users {
						if int64(id) == p.ID {
							if usernameOK {
								users[i].Username = username
							}
							if mailOK {
								users[i].Mail = mail
							}
							if passwordOK {
								users[i].Password = password
							}
							user = users[i]
							break
						}
					}
					return user, nil
				},
			},

			/* Delete user by id
			http://localhost:8000/user?query=mutation+_{delete(id:1){id,username,mail,password}}
			*/
			"delete": &graphql.Field{
				Type:        userType,
				Description: "Delete user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					user := User{}
					for i, p := range users {
						if int64(id) == p.ID {
							user = users[i]
							users = append(users[:i], users[i+1:]...)
						}
					}

					return user, nil
				},
			},
		},
	},
)
