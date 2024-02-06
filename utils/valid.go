package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 获取结构体中Tag的msg参数
func GetValidMsg(err error, obj any) string {
	// 使用的时候,传入obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 获取字段名
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
