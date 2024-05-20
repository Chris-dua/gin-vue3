package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_name", imagesApi.ImageNameListView)
	router.POST("images", middleware.JwtAuth(), imagesApi.ImageUploadView)
	router.DELETE("images", middleware.JwtAuth(), imagesApi.ImageRemoveView)
	router.PUT("images", middleware.JwtAuth(), imagesApi.ImageUpdateView)
	router.POST("image", middleware.JwtAuth(), imagesApi.ImageUploadDataView)
}
