package provider

import (
	"fmt"
	"games/app/global"
	"github.com/go-playground/locales/zh"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

type validateProvider struct {
}

var ValidateProvider = &validateProvider{}

func (p *validateProvider) Register() {
	validate := validator.New()
	uni := unTrans.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}

	global.Validator = validate
}
