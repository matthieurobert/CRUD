package main

import (
	"github.com/graphql-go/graphql"
)

//User contains informations about on user
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

// Création de l'objet user dans graphql
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"mail": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
