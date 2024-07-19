package category

type ListResponse struct {
	Total       int64 `json:"total" example:"10"`
	PageSize    int64 `json:"pageSize" example:"10"`
	CurrentPage int64 `json:"currentPage" example:"1"`
	MaxPage     int64 `json:"maxPage" example:"10"`
	Data        int64 `json:"data" example:"10"`
}

type ListItemResponse struct {
	Title       string `json:"title" example:"单人"`
	Value       int64  `json:"value" example:"99"`
	CurrentPage int64  `json:"currentPage" example:"1"`
	MaxPage     int64  `json:"maxPage" example:"10"`
	Data        int64  `json:"data" example:"10"`
}
