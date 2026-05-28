package initializers

import (
	"golang.org/x/text/language"

	"github.com/gofiber/contrib/v3/i18n"
)

var TranslateConfig *i18n.Config

func SetUpTranslator() {
	TranslateConfig = &i18n.Config{
		RootPath:         "./src/localize",
		AcceptLanguages:  []language.Tag{language.Persian, language.English},
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
	}
}
