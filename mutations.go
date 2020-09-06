package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
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
					sqlStatement := `
					INSERT INTO users (username, mail, password) 
					VALUES ($1, $2, $3)`
					db := postgres()
					_, err := db.Exec(sqlStatement, params.Args["username"].(string), params.Args["mail"].(string), params.Args["password"].(string))
					if err != nil {
						panic(err)
					}

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
					db := postgres()
					id, _ := params.Args["id"].(int)
					username, usernameOK := params.Args["username"].(string)
					mail, mailOK := params.Args["mail"].(string)
					password, passwordOK := params.Args["password"].(string)
					user := User{}

					var idMax int
					db.QueryRow("SELECT MAX(id) FROM users").Scan(&idMax)

					fmt.Println(idMax)

					for i := 0; i <= idMax; i++ {
						if id == i {
							if usernameOK {
								sqlStatement := `
								UPDATE users
								SET username = $2
								WHERE id = $1;`
								_, err := db.Exec(sqlStatement, id, username)
								if err != nil {
									panic(err)
								}
							}
							if mailOK {
								sqlStatement := `
								UPDATE users
								SET mail = $2
								WHERE id = $1;`
								_, err := db.Exec(sqlStatement, id, mail)
								if err != nil {
									panic(err)
								}
							}
							if passwordOK {
								sqlStatement := `
								UPDATE users
								SET password = $2
								WHERE id = $1;`
								_, err := db.Exec(sqlStatement, id, password)
								if err != nil {
									panic(err)
								}
							}
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
					db := postgres()
					id, _ := params.Args["id"].(int)
					user := User{}

					var idMax int
					db.QueryRow("SELECT MAX(id) FROM users").Scan(&idMax)

					for i := 1; i <= idMax; i++ {
						if id == i {
							sqlStatement := `
							DELETE FROM users
							WHERE id = $1;`
							_, err := db.Exec(sqlStatement, i)
							if err != nil {
								panic(err)
							}
						}
					}

					return user, nil
				},
			},
		},
	},
)
