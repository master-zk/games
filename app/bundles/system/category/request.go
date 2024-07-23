package category

type DeleteRequest struct {
	Id int64 `json:"id" binding:"required,numeric"`
}

type RecommendRequest struct {
	Id int64 `form:"id" binding:"required,numeric"`
}
