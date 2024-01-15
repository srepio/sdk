package types

type Paginated[T any] struct {
	Count      int `json:"total"`
	Page       int `json:"current_page"`
	TotalPages int `json:"total_pages"`
	Data       []T `json:"data"`
}

func NewPaginated[T any](count, page, totalPages int, data []T) *Paginated[T] {
	return &Paginated[T]{
		Count:      count,
		Page:       page,
		TotalPages: totalPages,
		Data:       data,
	}
}
