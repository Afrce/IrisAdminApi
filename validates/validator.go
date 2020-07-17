package validates

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/translations/zh"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	err      error
	Trans    ut.Translator
)

func init() {
	zh_cn := zh.New()
	uni = ut.New(zh_cn)
	Trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()

	// 设置对应的 中文字段名称
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("comment")
	})
	// 验证器注册翻译器
	if err = zh_translations.RegisterDefaultTranslations(Validate, Trans); err != nil {
		panic(fmt.Sprintf("注册验证翻译器错误：", err.Error()))
	}
}
