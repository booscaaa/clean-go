package productgraphqlservice

import (
	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/graphql-go/graphql"
)

func (service service) Fetch(params graphql.ResolveParams) (interface{}, error) {
	paginationRequest, _ := dto.FromValuePaginationGraphRequestParams(params)

	products, err := service.usecase.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
