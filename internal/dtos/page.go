package dtos

type Pagination struct {
	Page int    `json:"page"`
	Size int    `json:"size"`
	Sort string `json:"sort"`
}

type PaginatedList[T any] struct {
	TotalPages      int  `json:"totalPages"`
	PageSize        int  `json:"pageSize"`
	CurrentPage     int  `json:"currentPage"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	ItemsCount      int  `json:"itemsCount"`
	Content         []T  `json:"content"`
}

func NewPaginatedList[T any](items []T, page int, size int) *PaginatedList[T] {
	paginatedList := &PaginatedList[T]{
		PageSize:    size,
		CurrentPage: page,
		ItemsCount:  len(items),
	}
	paginatedList.Content = items[(page-1)*size : size*page]
	paginatedList.HasPreviousPage = page > 1
	paginatedList.TotalPages = (paginatedList.ItemsCount / size) + 1
	paginatedList.HasNextPage = page < paginatedList.TotalPages
	return paginatedList
}
