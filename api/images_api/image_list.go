package images_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	// 解析参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.CommonList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
	return
}
