package routers

import "server/api"

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menu_create", menuApi.MenuCreateView)
	router.GET("menus", menuApi.MenuListView)
	router.GET("menu_names", menuApi.MenuNameList)
	router.POST("menu_update", menuApi.MenuUpdateView)
	router.POST("menu_remove", menuApi.MenuRemoveView)
	router.GET("menu/:id", menuApi.MenuDetails)
}
