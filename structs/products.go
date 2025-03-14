package structs

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/gorm"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

type Product struct {
	gorm.Model
	ProductName string `json:"productName" gorm:"not null" validate:"required"`
	SKU         string `json:"sku" gorm:"not null" validate:"required"`
	Quantity    int    `json:"quantity" gorm:"not null" validate:"required"`
	Status      string `json:"status" gorm:"not null;default:tersedia" validate:"required"`
}

func init() {
	validate = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	translator, _ = uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, translator)

	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

func (p *Product) Validate() error {
	err := validate.Struct(p)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Translate(translator))
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}
