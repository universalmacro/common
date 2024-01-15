package dao

type List[T any] struct {
	Items      []T        `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
	Index int64 `json:"index"`
}
