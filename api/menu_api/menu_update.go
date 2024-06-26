package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdateView(context *gin.Context) {
	var cr MenuRequest
	err := context.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, context)
		return
	}
	id := context.Param("id")

	// 先把之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", context)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 如果选择了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", context)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", context)
		return
	}

	res.OkWithMessage("修改菜单成功", context)

}
