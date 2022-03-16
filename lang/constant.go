package lang

import (
	"encoding/json"

	"golang.org/x/text/language"
)

const (
	defaultFormatBundleFile = "json"
	defaultRootPath         = "./lang"
)

var (
	defaultLanguage       = language.English
	defaultUnmarshalFunc  = json.Unmarshal
	defaultAcceptLanguage = []language.Tag{
		defaultLanguage,
		language.Chinese,
		language.Hindi,
	}

	defaultBundleConfig = &BundleCfg{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
	}
)
