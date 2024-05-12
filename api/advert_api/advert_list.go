package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"
)

// AdvertListView 广告列表
// @Tags  广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo	false	"查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(context *gin.Context) {
	var cr models.PageInfo
	if err := context.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	// 判断 Referer 是否包含admin，如果是，就全部返回，不是，就返回is_show=true
	referer := context.GetHeader("Referer")
	fmt.Println(referer)
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin来的
		isShow = false
	}
	fmt.Println(isShow)
	list, count, _ := common.ComListFind(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, context)
}
