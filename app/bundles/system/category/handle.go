package category

import (
	"errors"
	"fmt"
	"games/app/common"
	"games/app/dal/system/model"
	"games/app/global"
	"github.com/gin-gonic/gin"
)

func Recommend(ctx *gin.Context) {
	response, err := recommendProcess(ctx)
	if err != nil {
		common.Fail(ctx, err.Error())
	} else {
		common.Success(ctx, response)
	}
}

func recommendProcess(c *gin.Context) (ListResponse, error) {
	var response ListResponse
	var err error
	var request RecommendRequest
	if err = c.ShouldBindQuery(&request); err != nil {
		goto RecommendProcessEnd
	} else {
		var res *[]model.Category
		db := global.DB.Jenny.Model(model.Category{})
		db.Where("pid = ?", request.Id).
			Where("is_menu = ?", 1).
			Select("id,title,icon").
			Order("id asc").
			Limit(int(common.PageSize)).
			Find(&res)

		response.Total = int64(len(*res))
		response.PageSize = common.PageSize
		response.CurrentPage = common.Page
		response.MaxPage = common.Page
		if response.Total > 0 {
			categories := *res
			var item common.EnumIntItemResponse
			for _, value := range categories {
				item.Title = value.Title
				item.Icon = value.Icon
				item.Value = int64(value.ID)
				response.Data = append(response.Data, item)
			}
		}
	}

RecommendProcessEnd:
	return response, err
}

func Menu(ctx *gin.Context) {
	common.Success(ctx, menuProcess())
}

func menuProcess() ListResponse {
	var response ListResponse
	var res *[]model.Category
	db := global.DB.Jenny.Model(model.Category{})
	db.Where("pid = ?", 0).
		Where("is_menu = ?", 1).
		Select("id,title,icon").
		Order("id asc").
		Limit(int(common.PageSize)).
		Find(&res)

	response.Total = int64(len(*res))
	response.PageSize = common.PageSize
	response.CurrentPage = common.Page
	response.MaxPage = common.Page
	if response.Total > 0 {
		categories := *res
		var item common.EnumIntItemResponse
		for _, value := range categories {
			item.Title = value.Title
			item.Icon = value.Icon
			item.Value = int64(value.ID)
			response.Data = append(response.Data, item)
		}
	}

	return response
}

func Delete(ctx *gin.Context) {
	err := deleteProcess(ctx)
	if err != nil {
		common.Fail(ctx, err.Error())
	} else {
		common.Success(ctx, nil)
	}
}

func deleteProcess(c *gin.Context) error {
	var err error
	var request DeleteRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		goto DeletedProcessEnd
	} else {
		var res *model.Category
		db := global.DB.Jenny.Model(model.Category{})
		db.Where("id = ?", request.Id).First(&res)
		fmt.Println(res)
		if (*res).ID > 0 {
			db.Where("id = ?", request.Id).Delete(&res)
		} else {
			err = errors.New("目标不存在")
		}
	}

DeletedProcessEnd:
	return err
}
