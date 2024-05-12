package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ImageListView 图片列表
// @Tags  图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo	false	"查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
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
