package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/boooscaaa/clean-go/core/dto"
)

// @Summary Create new product
// @Description Create new product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductRequest true "product"
// @Success 200 {object} domain.Product
// @Router /product [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := service.usecase.Create(productRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
