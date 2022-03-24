package productservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boooscaaa/clean-go/adapter/http/productservice"
	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/domain/mocks"
	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
)

func setupCreate(t *testing.T) (dto.CreateProductRequest, domain.Product, *gomock.Controller) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakeProductRequest, fakeProduct, mockCtrl
}

func TestCreate(t *testing.T) {
	fakeProductRequest, fakeProduct, mock := setupCreate(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(&fakeProduct, nil)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupCreate(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_PorductError(t *testing.T) {
	fakeProductRequest, _, mock := setupCreate(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
