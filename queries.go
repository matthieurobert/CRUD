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
					db := postgres()
					id, ok := u.Args["id"].(int)
					if ok {

						var idMax int
						db.QueryRow("SELECT MAX(id) FROM users").Scan(&idMax)

						for i := 1; i <= idMax; i++ {
							if id == i {
								sqlStatement := `SELECT id, username, mail, password FROM users WHERE id=$1;`
								var id int64
								var username string
								var mail string
								var password string

								db.QueryRow(sqlStatement, i).Scan(&id, &username, &mail, &password)

								user := User{
									ID:       id,
									Username: username,
									Mail:     mail,
									Password: password,
								}

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
					db := postgres()
					var users = []User{}

					var idMax int
					db.QueryRow("SELECT MAX(id) FROM users").Scan(&idMax)

					for i := 1; i <= idMax; i++ {

						sqlStatement := `SELECT id, username, mail, password FROM users WHERE id=$1;`
						var id int64
						var username string
						var mail string
						var password string

						db.QueryRow(sqlStatement, i).Scan(&id, &username, &mail, &password)

						if id != 0 {
							user := User{
								ID:       id,
								Username: username,
								Mail:     mail,
								Password: password,
							}

							users = append(users, user)
						}

					}

					return users, nil
				},
			},
		},
	},
)
