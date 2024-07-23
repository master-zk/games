package category

import "games/app/common"

type ListResponse struct {
	Total       int64                        `json:"total" example:"10"`
	PageSize    int64                        `json:"pageSize" example:"10"`
	CurrentPage int64                        `json:"currentPage" example:"1"`
	MaxPage     int64                        `json:"maxPage" example:"10"`
	Data        []common.EnumIntItemResponse `json:"data" example:"null"`
}
