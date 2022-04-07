package di

import (
	"github.com/boooscaaa/clean-go/adapter/http/graphql/productgraphqlservice"
	"github.com/boooscaaa/clean-go/adapter/http/rest/productservice"
	"github.com/boooscaaa/clean-go/adapter/postgres"
	"github.com/boooscaaa/clean-go/adapter/postgres/productrepository"
	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}

// ConfigProductGraphQLDI return a PoductGraphQLService abstraction with dependency injection configuration
func ConfigProductGraphQLDI(conn postgres.PoolInterface) domain.PoductGraphQLService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productgraphqlservice.New(productUseCase)

	return productService
}
