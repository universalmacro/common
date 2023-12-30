package dao

type List[T any] struct {
	Items      []T        `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Limit int `json:"limit"`
	Total int `json:"total"`
	Index int `json:"index"`
}
