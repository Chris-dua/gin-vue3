package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(context *gin.Context) {

	_claims, _ := context.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var userListReq UserListRequest
	if err := context.ShouldBindQuery(&userListReq); err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	var users = make([]UserResponse, 0)
	list, count, _ := common.ComListFind(models.UserModel{
		Role: ctype.Role(userListReq.Role),
	}, common.Option{
		PageInfo: userListReq.PageInfo,
		Likes:    []string{"nick_name", "user_name"},
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
		//users = append(users, user)
	}

	res.OkWithList(users, count, context)
}
