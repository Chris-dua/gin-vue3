package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (TagApi) TagRemoveView(context *gin.Context) {
	var cr models.RemoveRequest
	err := context.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, context)
		return
	}
	if len(cr.IDList) == 0 {
		res.FailWithMessage("无效的请求，id_list 不能为空", context)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", context)
		return
	}
	// 标签下有文章,需要展示文章数量
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), context)

}
