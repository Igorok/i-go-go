package ordertype

import (
	"github.com/graphql-go/graphql"
)

var CartItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CartItem",
		Fields: graphql.Fields{
			"idPr": &graphql.Field{
				Type: graphql.String,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var OrderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"idCl": &graphql.Field{
				Type: graphql.String,
			},
			"idEst": &graphql.Field{
				Type: graphql.String,
			},
			"cart": &graphql.Field{
				Type: graphql.NewList(CartItemType),
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"dateCreate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"dateCustomer": &graphql.Field{
				Type: graphql.DateTime,
			},
			"comment": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
