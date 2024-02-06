package common

import (
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string // 模糊匹配的字段
}

// CommonList 分页查询
func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}

	DB = DB.Where(model)
	// 模糊匹配
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(column+" like ?", "%"+option.Key+"%")
			continue
		}
		DB.Or(column+" like ?", "%"+option.Key+"%")
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	// 重置一下query
	query := DB.Where(model)
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
