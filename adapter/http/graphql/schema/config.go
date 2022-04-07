package schema

import (
	"fmt"

	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/graphql-go/graphql"
)

func Config(productService domain.PoductGraphQLService) graphql.Schema {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type:        paginationProducts,
				Description: "Get all products with server pagination",
				Args:        paginationRequestParams,
				Resolve:     productService.Fetch,
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	return schema
}

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
