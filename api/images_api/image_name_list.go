package images_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type ImageNameResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameList 图片名称列表(不分页)
// @Tags 图片管理
// @Summary 图片名称列表(不分页)
// @Description 图片名称列表(不分页)
// @Router /api/image_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageNameResponse}
func (ImagesApi) ImageNameList(c *gin.Context) {
	var imageList []ImageNameResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
	return
}
