package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (ImagesApi) ImageRemoveView(context *gin.Context) {
	var cr models.RemoveRequest
	err := context.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", context)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), context)
}
