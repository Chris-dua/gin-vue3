package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_name", imagesApi.ImageNameListView)
	router.POST("images", imagesApi.ImageUploadView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	router.PUT("images", imagesApi.ImageUpdateView)
}
