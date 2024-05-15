package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (TagApi) TagUpdateView(context *gin.Context) {

	id := context.Param("id")
	var cr TagRequest
	err := context.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, context)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", context)
		return
	}
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", context)
		return
	}

	res.OkWithMessage("修改标签成功", context)
}
