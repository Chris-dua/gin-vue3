package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("email_login", userApi.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), userApi.LogoutView)
}
