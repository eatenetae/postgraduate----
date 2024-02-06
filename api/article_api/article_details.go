package article_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/es_ser"
	"server/service/redis_ser"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 增加浏览量
	redis_ser.NewLook().Set(cr.ID)

	model, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommonDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
