package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"server/api"
	"server/middleware"
)

var store = cookie.NewStore([]byte("secret"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("login", userApi.LoginView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.POST("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.POST("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), userApi.LogoutView)
	router.POST("user_remove", middleware.JwtAdmin(), userApi.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
	router.POST("user_create", middleware.JwtAdmin(), userApi.UserCreateView)
}
