package routers

import "server/api"

func (router RouterGroup) LogRouter() {
	logApi := api.ApiGroupApp.LogApi
	router.POST("log_list", logApi.LogListView)
	router.POST("log_remove", logApi.LogRemoveListView)
}
