package schema

import "github.com/graphql-go/graphql"

var paginationRequestParams = graphql.FieldConfigArgument{
	"page": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"itemsPerPage": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"search": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"descending": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"sort": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
