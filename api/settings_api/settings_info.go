package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}

}
