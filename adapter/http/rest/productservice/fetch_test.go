package productservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boooscaaa/clean-go/adapter/http/rest/productservice"
	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/domain/mocks"
	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
)

func setupFetch(t *testing.T) (dto.PaginationRequestParms, domain.Product, *gomock.Controller) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeProduct, mockCtrl
}

func TestFetch(t *testing.T) {
	fakePaginationRequestParams, fakeProduct, mock := setupFetch(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.Product{fakeProduct},
		Total: 1,
	}, nil)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetch_PorductError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupFetch(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
