package main

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	// 创建一个验证器实例
	validate = validator.New()

	registerZhTrans()

	// 自定义字段名的显示
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("zhtrans")
		if name == "" {
			return fld.Name
		}
		return name
	})
}

// 注册一个中文翻译器
func registerZhTrans() {
	zhTrans := zh.New()
	uni := ut.New(zhTrans, zhTrans)
	trans, _ = uni.GetTranslator("zh")
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(err)
	}
}

func Validate(data interface{}) error {
	return validateV1(data)
}

func validateV1(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		errText := make([]string, 0)
		for _, vaErr := range err.(validator.ValidationErrors) {
			errMsg := vaErr.Translate(trans)
			errText = append(errText, errMsg)
		}
		err = errors.New(strings.Join(errText, ","))
	}
	return err
}
