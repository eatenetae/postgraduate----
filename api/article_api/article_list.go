package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"server/global"
	"server/models"
	"server/models/res"
	"server/service/es_ser"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommonList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	// 排除 不显示list特点 的字段, 例如不显示content字段
	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
