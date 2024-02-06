package flag

import (
	"server/global"
	"server/models"
	"server/plugins/log_stash"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
			&models.UserCollectModel{},
			&log_stash.LogStashModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
