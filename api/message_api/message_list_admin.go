package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (MessageApi) MessageListAllView(context *gin.Context) {
	var cr models.PageInfo
	if err := context.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	list, count, _ := common.ComListFind(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, context)
}
