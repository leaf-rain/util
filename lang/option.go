package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	// Option ...
	Option func(I18n)
)

// BundleCfg ...
type BundleCfg struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
}

// WithBundle ...
func WithBundle(config *BundleCfg) Option {
	return func(g I18n) {
		g.setBundle(config)
	}
}
