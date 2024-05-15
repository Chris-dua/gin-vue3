package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.POST("tags", tagApi.TagCreateView)
	router.GET("tags", tagApi.TagListView)
	router.PUT("tags/:id", tagApi.TagUpdateView)
	router.DELETE("tags", tagApi.TagRemoveView)

}
