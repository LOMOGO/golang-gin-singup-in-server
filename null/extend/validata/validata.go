//用于验证接口值的合法性
package validata

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Trans ut.Translator

func Setup() {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_trans.RegisterDefaultTranslations(v, Trans)
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("label")
			return name
		})
	}
}
