package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"xiaoyun/backend/utils"
)

// GetErrMsg 从验证错误中提取错误原因和对应的字段
func GetErrMsg(data interface{}, err error) string {
	// 检查错误是否为 validator.ValidationErrors 类型
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// 获取数据的反射类型
		dataType := reflect.TypeOf(data)
		if dataType.Kind() == reflect.Ptr {
			dataType = dataType.Elem() // 解引用指针类型
		}
		// 遍历所有验证错误
		for _, e := range validationErrors {
			// 获取字段名
			fieldName := e.Field()
			// 检查结构体中是否存在该字段
			if field, exists := dataType.FieldByName(fieldName); exists {
				// 获取字段的 "msg" 标签，如果不存在则使用默认错误消息
				msg := field.Tag.Get("msg")
				if msg == "" {
					msg = fmt.Sprintf("'%s'格式不正确", fieldName)
				}
				return msg
			}
		}
	}
	// 如果不是验证错误，返回通用错误消息
	return "Invalid input data"
}

func SetupValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("phone", phone)
		if err != nil {
			fmt.Println("phone validate register success")
		}
		err = v.RegisterValidation("username", username)
		if err != nil {
			fmt.Println("username validate register success")
		}
	}
}

func phone(fl validator.FieldLevel) bool {
	if data, ok := fl.Field().Interface().(string); ok {
		return utils.IsPhoneValid(data)
	}
	return false
}

func username(fl validator.FieldLevel) bool {
	if data, ok := fl.Field().Interface().(string); ok {
		return utils.IsAlphanumericValid(data)
	}
	return false
}
