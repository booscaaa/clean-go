package schema

import (
	"github.com/graphql-go/graphql"
)

var product = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var paginationProducts = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pagination",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type: graphql.NewList(product),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
