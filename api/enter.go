package api

import (
	"server/api/advert_api"
	"server/api/article_api"
	"server/api/data_api"
	"server/api/images_api"
	"server/api/log_api"
	"server/api/menu_api"
	"server/api/news_api"
	"server/api/settings_api"
	"server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	ArticleApi  article_api.ArticleApi
	NewsApi     news_api.NewsApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
}

var ApiGroupApp = new(ApiGroup)
