package dto

import (
	"net/http"
	"strconv"
	"strings"
)

// PaginationRequestParms is an representation query string params to filter and paginate products
type PaginationRequestParms struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

// FromValuePaginationRequestParams converts query string params to a PaginationRequestParms struct
func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParms, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	paginationRequestParms := PaginationRequestParms{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}

	return &paginationRequestParms, nil
}
