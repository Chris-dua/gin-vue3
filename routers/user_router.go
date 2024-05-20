package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gvb_server/api"
	"gvb_server/middleware"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", userApi.EmailLoginView)
	router.POST("users", middleware.JwtAdmin(), userApi.UserCreateView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.GET("user_info", middleware.JwtAuth(), userApi.UserInfoView)
	router.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), userApi.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), userApi.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
