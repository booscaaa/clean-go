package domain

type Pagination[T any] struct {
	Items T     `json:"items"`
	Total int32 `json:"total"`
}
