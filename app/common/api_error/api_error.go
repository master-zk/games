package api_error

import (
	"errors"
	"games/app/common"
	"gorm.io/gorm"
	"strconv"
)

type ApiError struct {
	Msg    string
	Status int
}

func (e *ApiError) Error() string {
	return e.Msg
}

func HandleError(err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		PanicApiError(err.Error(), common.CodeMessageError)
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		PanicApiError(err.Error(), common.CodeMessageError)
	} else if _, ok := err.(*strconv.NumError); ok {
		PanicApiError(err.Error(), common.CodeMessageError)
	} else if err != nil {
		panic(err)
	}
}

func PanicApiError(msg string, statuses ...int) {
	status := common.CodeMessageError
	if len(statuses) > 0 {
		status = statuses[0]
	}

	panic(&ApiError{msg, status})
}
