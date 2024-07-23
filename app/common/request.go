package common

const (
	PageSize int64 = 10
	Page     int64 = 1
)

type BasicRequest struct {
	page     int64
	pageSize int64
}
