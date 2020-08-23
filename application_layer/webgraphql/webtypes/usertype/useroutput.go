package usertype

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},

			"login": &graphql.Field{
				Type: graphql.String,
			},
			/*
				"password": &graphql.Field{
					Type: graphql.String,
				},
			*/
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phoneNumber": &graphql.Field{
				Type: graphql.String,
			},

			"name": &graphql.Field{
				Type: graphql.String,
			},
			"surname": &graphql.Field{
				Type: graphql.String,
			},
			"patronymic": &graphql.Field{
				Type: graphql.String,
			},

			"roles": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var SignInType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SignIn",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
