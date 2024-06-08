// Package validate contains the support for validating models.
package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en" 
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// 校验 validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

// 翻译 translator is a cache of locale and translation information.
var translator ut.Translator

func init() {

	// Instantiate a validator.
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Create a translator for english so the error messages are
	// more human-readable than technical.
	translator, _ = ut.New(en.New(), en.New()).GetTranslator("en")

	// Register the english error messages for use.
	en_translations.RegisterDefaultTranslations(validate, translator)

	// Use JSON tag names for errors instead of Go struct names.
	// 注册了一个标签名称函数，用于将结构体字段的 JSON 标签名称作为验证错误消息中的字段名称，
	// 而不是使用 Go 结构体的字段名称。
	// 这个函数通过解析结构体字段的 json 标签来获取字段的 JSON 名称，如果标签为 -，则返回空字符串。
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Check validates the provided model against it's declared tags.
func Check(val any) error {
	if err := validate.Struct(val); err != nil {
		verrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		var fields FieldErrors
		for _, verror := range verrors {
			fields = append(fields, FieldError{
				Field: verror.Field(),
				Err:   verror.Translate(translator),
			})
		}

		return fields
	}

	return nil
}