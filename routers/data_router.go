package routers

import "server/api"

func (router RouterGroup) DataRouter() {
	dataApi := api.ApiGroupApp.DataApi
	router.GET("data_week_login", dataApi.SevenLogin)
	router.GET("data_sum", dataApi.DataSumView)
}
