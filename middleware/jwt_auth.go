package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", context)
			context.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", context)
			context.Abort()
			return
		}
		// 判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", context)
			context.Abort()
			return
		}
		// 登录的用户
		context.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", context)
			context.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", context)
			context.Abort()
			return
		}
		// 判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", context)
			context.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", context)
			context.Abort()
			return
		}
		context.Set("claims", claims)
	}
}