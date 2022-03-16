package lang

import (
	"encoding/json"

	"golang.org/x/text/language"
)

var atI18n I18n

// newI18n ...
func newI18n(opts ...Option) {
	// init default value
	ins := &i18nImpl{}
	ins.setBundle(defaultBundleConfig)

	// overwrite default value by options
	for _, opt := range opts {
		opt(ins)
	}

	atI18n = ins
}

func Init(langPath string, defaultLang language.Tag, accepts ...language.Tag) I18n {
	if len(accepts) == 0 {
		accepts = []language.Tag{language.English, language.Hindi, language.Chinese}
	}
	opts := WithBundle(&BundleCfg{
		RootPath:         langPath,
		AcceptLanguage:   accepts,
		DefaultLanguage:  defaultLang,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})
	newI18n(opts)
	return atI18n
}

/*GetMessage get the i18n message
 param is one of these type: messageID, *i18n.LocalizeConfig
 Example:
	GetMessage("hello") // messageID is hello
	GetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
	})
*/
func GetMessage(param interface{}, lng ...string) (string, error) {
	return atI18n.getMessage(param, lng...)
}

/*MustGetMessage get the i18n message without error handling
  param is one of these type: messageID, *i18n.LocalizeConfig
  Example:
	MustGetMessage("hello") // messageID is hello
	MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
	})
*/
func MustGetMessage(param interface{}, lng ...string) string {
	return atI18n.mustGetMessage(param, lng...)
}

func GetIds() []string {
	return atI18n.getIds()
}
