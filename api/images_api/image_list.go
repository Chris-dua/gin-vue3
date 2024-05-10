package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (ImagesApi) ImageListView(context *gin.Context) {
	var cr models.PageInfo
	//查询页数
	err := context.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	list, count, err := common.ComListFind(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OkWithList(list, count, context)
	return
}
