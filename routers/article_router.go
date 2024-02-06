package routers

import (
	"server/api"
	"server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("article_create", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("articles", articleApi.ArticleListView)
	router.GET("article/detail", articleApi.ArticleDetailView)
	router.GET("article/tags", articleApi.ArticleTagListView)
	router.POST("article_update", articleApi.ArticleUpdateView)
	router.GET("article/:id", articleApi.ArticleDetailView)
	router.POST("article_remove", articleApi.ArticleRemoveView)
	router.POST("article_collect", middleware.JwtAuth(), articleApi.ArticleCollCreateView)
	router.POST("article/full_text_search", articleApi.FullTextSearchView)
}
