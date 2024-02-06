package routers

import (
	"server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.POST("settings/:name", settingsApi.SettingsInfoUpdateView)
}
