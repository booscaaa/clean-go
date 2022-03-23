package productusecase

import (
	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
