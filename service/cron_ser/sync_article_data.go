package cron_ser

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"server/global"
	"server/models"
	"server/service/redis_ser"
)

func SyncArticleData() {
	// 1. 查询es中的全部数据, 为后面的数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	// 2.拿到redis中的缓存数据
	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err := json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		// 拿到缓存数据
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3. 更新数据
		newDigg := digg + article.DiggCount
		newLook := look + article.LookCount
		newComment := comment + article.CommentCount

		// 4. 判断是否有变化, 如果没变化的话
		if digg == 0 && look == 0 && comment == 0 {
			global.Log.Infof("%s无变化", article.Title)
			continue
		}
		// 5. 保存到es
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 点赞数:%d, 评论数:%d, 浏览量:%d", article.Title, newComment, newLook)
	}
	//6. 清除redis
	redis_ser.NewDigg().Clear()
	redis_ser.NewLook().Clear()
	redis_ser.NewCommentCount().Clear()
}
