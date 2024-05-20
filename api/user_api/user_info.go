package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

func (UserApi) UserInfoView(context *gin.Context) {

	_claims, _ := context.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var userModel models.UserModel
	err := global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", context)
		return
	}

	res.OkWithData(userModel, context)
}
