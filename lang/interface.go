package lang

type I18n interface {
	getMessage(param interface{}, lng ...string) (string, error)
	mustGetMessage(param interface{}, lng ...string) string
	setBundle(cfg *BundleCfg)
	getIds() []string
}
