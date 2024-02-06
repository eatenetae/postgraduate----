package routers

import "server/api"

func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	router.POST("advert_create", advertApi.AdvertCreateView)
	router.GET("adverts", advertApi.AdvertListView)
	router.POST("advert_update/:id", advertApi.AdvertUpdateView)
	router.POST("advert_remove", advertApi.AdvertRemoveView)
}
