package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {

	var bundle = &i18n.Bundle{}
	bundle.MustLoadMessageFile("./i18n/en.json")
	bundle.MustLoadMessageFile("./i18n/el.json")

	loc := i18n.NewLocalizer(bundle, "en")

	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "messages",
		TemplateData: map[string]interface{}{
			"Name":  "Alex",
			"Count": 10,
		},
		PluralCount: 10,
	})

	fmt.Println(translation)

}
