package lang

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ I18n = (*i18nImpl)(nil)

type i18nImpl struct {
	bundle          *i18n.Bundle
	localizerByLng  map[string]*i18n.Localizer
	defaultLanguage language.Tag
	ids             []string
}

// getMessage get lang message by lng and messageID
func (i *i18nImpl) getMessage(param interface{}, lng ...string) (string, error) {
	lngTag := i.defaultLanguage.String()
	if len(lng) > 0 {
		lngTag = lng[0]
	}
	localizer := i.getLocalizerByLng(lngTag)

	localizeConfig, err := i.getLocalizeConfig(param)
	if err != nil {
		return "", err
	}
	message, err := localizer.Localize(localizeConfig)
	if err != nil {
		return "", err
	}

	return message, nil
}

// mustGetMessage ...
func (i *i18nImpl) mustGetMessage(param interface{}, lng ...string) string {
	lngTag := i.defaultLanguage.String()
	if len(lng) > 0 {
		lngTag = lng[0]
	}
	message, _ := i.getMessage(param, lngTag)
	return message
}

func (i *i18nImpl) setBundle(cfg *BundleCfg) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(cfg.FormatBundleFile, cfg.UnmarshalFunc)

	i.bundle = bundle
	i.defaultLanguage = cfg.DefaultLanguage

	i.loadMessageFiles(cfg)
	i.setLocalizerByLng(cfg.AcceptLanguage)
}

// loadMessageFiles load all file lang to bundle
func (i *i18nImpl) loadMessageFiles(config *BundleCfg) {
	var msgMap = make(map[string]struct{})
	for _, lng := range config.AcceptLanguage {
		// path := fmt.Sprintf("%s/%s.%s", config.RootPath, lng.String(), config.FormatBundleFile)
		path := config.RootPath + "/" + lng.String() + "." + config.FormatBundleFile
		path, _ = filepath.Abs(path)
		// i.bundle.MustLoadMessageFile(path)
		msgFile, err := i.bundle.LoadMessageFile(path)
		if err != nil || msgFile == nil {
			fmt.Printf("[lang i18n] load file error. %v\n", err)
			continue
		}
		for index := range msgFile.Messages {
			if _, ok := msgMap[msgFile.Messages[index].ID]; !ok {
				msgMap[msgFile.Messages[index].ID] = struct{}{}
				if len(i.ids) == 0 {
					i.ids = make([]string, 0)
				}
				i.ids = append(i.ids, msgFile.Messages[index].ID)
			}
		}
	}
}

// setLocalizerByLng set localizer by language
func (i *i18nImpl) setLocalizerByLng(acceptLanguage []language.Tag) {
	i.localizerByLng = map[string]*i18n.Localizer{}
	for _, lng := range acceptLanguage {
		lngStr := lng.String()
		i.localizerByLng[lngStr] = i.newLocalizer(lngStr)
	}

	// set defaultLanguage if it isn't exist
	defaultLng := i.defaultLanguage.String()
	if _, hasDefaultLng := i.localizerByLng[defaultLng]; !hasDefaultLng {
		i.localizerByLng[defaultLng] = i.newLocalizer(defaultLng)
	}
}

// newLocalizer create a localizer by language
func (i *i18nImpl) newLocalizer(lng string) *i18n.Localizer {
	lngDefault := i.defaultLanguage.String()
	lngs := []string{
		lng,
	}

	if lng != lngDefault {
		lngs = append(lngs, lngDefault)
	}

	localizer := i18n.NewLocalizer(
		i.bundle,
		lngs...,
	)
	return localizer
}

// getLocalizerByLng get localizer by language
func (i *i18nImpl) getLocalizerByLng(lng string) *i18n.Localizer {
	localizer, hasValue := i.localizerByLng[lng]
	if hasValue {
		return localizer
	}

	return i.localizerByLng[i.defaultLanguage.String()]
}

func (i *i18nImpl) getLocalizeConfig(param interface{}) (*i18n.LocalizeConfig, error) {
	switch paramValue := param.(type) {
	case string:
		localizeConfig := &i18n.LocalizeConfig{
			MessageID: paramValue,
		}
		return localizeConfig, nil
	case *i18n.LocalizeConfig:
		return paramValue, nil
	}

	msg := fmt.Sprintf("un supported lang param: %v", param)
	return nil, errors.New(msg)
}

func (i *i18nImpl) getIds() []string {
	return i.ids
}
