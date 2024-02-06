package routers

import (
	"server/api"
)

func (router RouterGroup) NewsRouter() {
	newsApi := api.ApiGroupApp.NewsApi
	router.POST("news", newsApi.NewsListView)

}
