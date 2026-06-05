package initializers

import (
	"golang.org/x/text/language"

	"github.com/gofiber/contrib/v3/i18n"
)

var Translator *i18n.I18n

func SetUpTranslator() {
	Translator = i18n.New(
		&i18n.Config{
			RootPath: "./src/localize",
			AcceptLanguages: []language.Tag{
				language.Persian,
				language.English,
			},
			DefaultLanguage:  language.English,
			FormatBundleFile: "json",
		},
	)
}
