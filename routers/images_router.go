package routers

import "server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("image_upload", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("image_names", imagesApi.ImageNameList)
	router.POST("remove_images", imagesApi.ImageRemoveView)
	router.POST("update_image", imagesApi.ImageUpdateView)
}
