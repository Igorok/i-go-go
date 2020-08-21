package ordertype

import (
	"github.com/graphql-go/graphql"
)

var CartItemArgType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CartItemArgType",
		Fields: graphql.InputObjectConfigFieldMap{
			"idPr": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"count": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)
