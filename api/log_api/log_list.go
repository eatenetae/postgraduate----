package log_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/plugins/log_stash"
	"server/service/common"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.CommonList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
