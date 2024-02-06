package utils

import (
	"reflect"
	"strings"
)

func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	// 如果类型是指针,则去掉指针
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			// 只关注导出字段,
			if fieldType.IsExported() {
				// 排除掉gorm:"-"的字段
				if fieldType.Tag.Get("gorm") != "-" {
					continue
				}
				// 如果没有gorm:"column:xxx"的tag,则使用驼峰转下划线的规则
				name := Camel2Underscore(fieldType.Name)
				// gorm:"column:xxx"
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					// 截去column:前缀
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						// 如果后面还有分号,则截取前面的部分
						pos := strings.Index(content, ";")
						if pos > 0 {
							name = content[0:pos]
						} else if pos < 0 {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	}
	return nil
}

func Camel2Underscore(name string) string {
	buffer := strings.Builder{}
	for i, v := range name {
		if i == 0 {
			buffer.WriteRune(v)
		} else {
			if v >= 'A' && v <= 'Z' {
				buffer.WriteRune('_')
				buffer.WriteRune(v + 32)
			} else {
				buffer.WriteRune(v)
			}
		}
	}
	return buffer.String()
}

// GetGormFields 用法
//type User struct {
//	Id        int    `gorm:"column:id;primary_key"`
//	Name      string `gorm:"column:name"`
//	Age       int    `gorm:"column:age"`
//	CreatedAt int64  `gorm:"column:created_at"`
//	UpdatedAt int64  `gorm:"column:updated_at"`
//}
//var (
//	_all_user_field = GetGormFields(User{})
//)
//func GetUserByName(name string) *User {
//	db := GetDB()
//	var user User
//	if err := db.Select(_all_user_field...).Where("name =?", name).First(&user).Error; err!= nil {
//		return nil
//	} else {
//		return &user
//	}
//}
