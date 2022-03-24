package domain

// Pagination is representation of Fetch methods returns
type Pagination struct {
	Items interface{} `json:"items"`
	Total int32       `json:"total"`
}
