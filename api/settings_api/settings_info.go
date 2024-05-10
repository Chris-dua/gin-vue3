package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingApi) SettingsInfoView(context *gin.Context) {
	var cr SettingsUri
	err := context.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, context)
	case "email":
		res.OkWithData(global.Config.Email, context)
	case "qq":
		res.OkWithData(global.Config.QQ, context)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, context)
	case "jwt":
		res.OkWithData(global.Config.Jwt, context)
	default:
		res.FailWithMessage("没有对应的配置信息", context)
	}
}
