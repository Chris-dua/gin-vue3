package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), messageApi.MessageCreateView)
	router.GET("messages_all", middleware.JwtAdmin(), messageApi.MessageListAllView)
	router.GET("messages", middleware.JwtAuth(), messageApi.MessageListView)
	router.GET("message_record", middleware.JwtAuth(), messageApi.MessageRecordView)

}
