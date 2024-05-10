package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingApi) SettingsInfoUpdate(context *gin.Context) {
	var cr SettingsUri
	err := context.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, context)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := context.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, context)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err := context.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, context)
			return
		}
		global.Config.Email = info

	case "qq":
		var info config.QQ
		err := context.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, context)
			return
		}
		global.Config.QQ = info

	case "qiniu":
		var info config.QiNiu
		err := context.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, context)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err := context.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, context)
			return
		}
		global.Config.Jwt = info

	default:
		res.FailWithMessage("没有对应的配置信息", context)
	}
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), context)
		return
	}
	res.OkWithout(context)

}
