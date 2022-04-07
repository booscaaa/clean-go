package productgraphqlservice

import "github.com/boooscaaa/clean-go/core/domain"

type service struct {
	usecase domain.ProductUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.ProductUseCase) domain.PoductGraphQLService {
	return &service{
		usecase: usecase,
	}
}
