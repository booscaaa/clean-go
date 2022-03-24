package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/domain/mocks"
	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/boooscaaa/clean-go/core/usecase/productusecase"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)
}

func TestCreate_Error(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	faker.FakeData(&fakeRequestProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.NotNil(t, err)
	require.Nil(t, product)
}
