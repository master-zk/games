package response

type ListResponse struct {
	Total       int64 `json:"total" example:"99"`
	PageSize    int64 `json:"pageSize" example:"10"`
	CurrentPage int64 `json:"currentPage" example:"1"`
	MaxPage     int64 `json:"maxPage" example:"10"`
	Data        int64 `json:"data" example:"10"`
}

type EnumIntItemResponse struct {
	Title string `json:"title" example:"int枚举1"`
	Value int64  `json:"value" example:"1"`
}

type EnumStringItemResponse struct {
	Title string `json:"title" example:"string枚举1"`
	Value string `json:"value" example:"code"`
}
