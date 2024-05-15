package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("articles", articleApi.ArticleListView)
	router.GET("articles/detail", articleApi.ArticleDetailByTitleView)
	router.GET("articles/calendar", articleApi.ArticleCalendarView)
	router.GET("articles/tags", articleApi.ArticleTagListView)
	router.PUT("articles", articleApi.ArticleUpdateView)
	router.DELETE("articles", articleApi.ArticleRemoveView)
	router.GET("articles/:id", articleApi.ArticleDetailView)
}
