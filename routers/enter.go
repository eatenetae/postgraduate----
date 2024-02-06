package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"server/global"
)

type RouterGroup struct {
	*gin.RouterGroup //匿名字段, 可以直接通过 结构体的名字 调用此字段的方法
}

func InitRouter() *gin.Engine {
	// 设置gin模式 debug/release, 这里设置release模式
	gin.SetMode(global.Config.System.Env)
	// 初始化路由, 创建了一个gin引擎实例
	router := gin.Default()
	// 为swagger配置路由
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 分配路由组 "/api"
	apiRouterGroup := router.Group("/api")
	// 实例化路由组
	routerGroupApp := RouterGroup{apiRouterGroup}

	// 调用路由组的方法,为settings模块注册路由
	routerGroupApp.SettingsRouter()
	// 调用路由组的方法,为images模块注册路由
	routerGroupApp.ImagesRouter()
	// 调用路由组的方法,为advert模块注册路由
	routerGroupApp.AdvertRouter()
	// 调用路由组的方法,为menu模块注册路由
	routerGroupApp.MenuRouter()
	// 调用路由组的方法,为user模块注册路由
	routerGroupApp.UserRouter()
	// 调用路由组的方法,为article模块注册路由
	routerGroupApp.ArticleRouter()
	// 调用路由组的方法,为news模块注册路由
	routerGroupApp.NewsRouter()

	// 调用路由组的方法,为log模块注册路由
	routerGroupApp.LogRouter()

	return router
}
