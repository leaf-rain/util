package logger

type ILoggerOption interface {
	Apply(*Option)
}

type Option struct {
	AppName   string
	Dir       string
	Level     string
	Formatter string
	ShowLine  bool
}

type appNameOption string

func (s appNameOption) Apply(o *Option) {
	o.AppName = string(s)
}

type dirOption string

func (s dirOption) Apply(o *Option) {
	o.Dir = string(s)
}

type levelOption string

func (s levelOption) Apply(o *Option) {
	o.Level = string(s)
}

type formatterOption string

func (f formatterOption) Apply(o *Option) {
	o.Formatter = string(f)
}

type showLineOption bool

func (s showLineOption) Apply(o *Option) {
	o.ShowLine = bool(s)
}

var defaultOptions = []ILoggerOption{appNameOption("app"), dirOption("logs"), levelOption("info"), formatterOption("text")}

func WithLogDir(dir string) dirOption {
	return dirOption(dir)
}

func WithLogLevel(level string) levelOption {
	return levelOption(level)
}

func WithLogAppName(appname string) appNameOption {
	return appNameOption(appname)
}

func WithLogFormatter(formatter string) formatterOption {
	return formatterOption(formatter)
}

func WithLogShowLine(s bool) showLineOption {
	return showLineOption(s)
}
