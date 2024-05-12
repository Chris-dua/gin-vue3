package routers

import "gvb_server/api"

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("email_login", userApi.EmailLoginView)

}
