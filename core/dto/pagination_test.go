package dto_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boooscaaa/clean-go/core/dto"

	"github.com/stretchr/testify/require"
)

func TestFromValuePaginationRequestParams(t *testing.T) {
	fakeRequest := httptest.NewRequest(http.MethodGet, "/product", nil)
	queryStringParams := fakeRequest.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	fakeRequest.URL.RawQuery = queryStringParams.Encode()

	paginationRequest, err := dto.FromValuePaginationRequestParams(fakeRequest)

	require.Nil(t, err)
	require.Equal(t, paginationRequest.Page, 1)
	require.Equal(t, paginationRequest.ItemsPerPage, 10)
	require.Equal(t, paginationRequest.Sort, []string{""})
	require.Equal(t, paginationRequest.Descending, []string{""})
	require.Equal(t, paginationRequest.Search, "")
}
