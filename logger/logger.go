package logger

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const FileSuffix = ".log"

type Logger struct {
	logrus.Logger
	o    Option
	sets []ILoggerOption
}

func New(opts ...ILoggerOption) *Logger {
	log := &Logger{
		o:    Option{},
		sets: append(defaultOptions, opts...),
	}
	for _, v := range log.sets {
		v.Apply(&log.o)
	}
	switch log.o.Formatter {
	case "text":
		log.Formatter = new(logrus.TextFormatter)
	case "json":
		log.Formatter = new(logrus.JSONFormatter)
	case "game":
		log.Formatter = NewGameFormater(log.o.AppName, TYPE_UNKNOW)
	default:
		panic(fmt.Sprintf("format %s not support.", log.o.Formatter))
	}
	lv, err := logrus.ParseLevel(log.o.Level)
	if err != nil {
		lv = logrus.InfoLevel
	}
	log.Hooks = make(logrus.LevelHooks)
	log.Out = os.Stderr
	log.Level = lv
	log.ExitFunc = os.Exit

	customLoger(&log.Logger, log.o)
	return log
}

func (log *Logger) HackDefaultLogger() {
	customLoger(logrus.StandardLogger(), log.o)
}

func customLoger(log *logrus.Logger, opt Option) {
	appName, logLevel, logFormater, showLine := opt.AppName, opt.Level, opt.Formatter, opt.ShowLine
	dir, _ := filepath.Abs(opt.Dir)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(level)
	}

	formatter := NewGameFormater(appName, TYPE_UNKNOW)

	// if len(logDebugTypes) > 0 {
	// 	for _, v := range logDebugTypes {
	// 		l := ParseGameLogType(v)
	// 		formater.AddType(l)
	// 	}
	// } else {
	// 	formater.AddType(TYPE_CGI | TYPE_GAME | TYPE_REDIS | TYPE_TCP)
	// }

	if showLine {
		log.SetReportCaller(true)
	}

	if dir != "" {
		os.MkdirAll(dir, os.ModePerm)
	}

	var fformater logrus.Formatter
	switch logFormater {
	case FORMAT_JSON:
		fformater = &logrus.JSONFormatter{CallerPrettyfier: callFormator}
	case FORMAT_TEXT:
		fformater = &logrus.TextFormatter{CallerPrettyfier: callFormator}
	case FORMAT_CUSTOM:
		fformater = formatter
	default:
		fformater = &logrus.TextFormatter{CallerPrettyfier: callFormator}
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: writer(dir, logrus.TraceLevel.String(), 7),
		logrus.DebugLevel: writer(dir, logrus.DebugLevel.String(), 7), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer(dir, logrus.InfoLevel.String(), 7),
		logrus.WarnLevel:  writer(dir, logrus.WarnLevel.String(), 7),
		logrus.ErrorLevel: writer(dir, logrus.ErrorLevel.String(), 7),
		logrus.FatalLevel: writer(dir, logrus.FatalLevel.String(), 7),
		logrus.PanicLevel: writer(dir, logrus.PanicLevel.String(), 7),
	}, fformater)
	log.AddHook(lfHook)
	// log.SetFormatter(fformater)

	splitOutter := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: os.Stdout, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  os.Stdout,
		logrus.WarnLevel:  os.Stderr,
		logrus.ErrorLevel: os.Stderr,
		logrus.FatalLevel: os.Stderr,
		logrus.PanicLevel: os.Stderr,
	}, formatter)
	log.AddHook(splitOutter)
	setNull(log)
}

func (l *Logger) Flush() {

}

func (l *Logger) SetLogDir(logDir string) {
	l.o.Dir = logDir
	if logDir != "" {
		os.MkdirAll(logDir, os.ModePerm)
	}
}
func (l *Logger) SetLogLevel(level string) {
	lv, _ := logrus.ParseLevel(level)
	l.SetLevel(lv)
}

func setNull(log *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func callFormator(f *runtime.Frame) (function string, file string) {
	//处理文件名
	fileName := path.Base(f.File)
	return f.Function, fmt.Sprintf("[%s:%d] ", fileName, f.Line)
}

func getLogFilename(dir string, lv logrus.Level) (logPath, linkName string) {
	hostname, _ := os.Hostname()
	processName := path.Base(os.Args[0])
	user, _ := user.Current()

	logPath = fmt.Sprintf("%s.%s.%s.log.%s", processName, hostname, user.Username, strings.ToUpper(lv.String()))
	linkName = fmt.Sprintf("%s.%s", processName, strings.ToUpper(lv.String()))

	logPath = path.Join(dir, logPath)
	linkName = path.Join(dir, linkName)
	return
}

func writer(logPath string, level string, save uint) *rotatelogs.RotateLogs {
	logFullPath, linkPath := getLogName(logPath, level)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fileSuffix := time.Now().In(cstSh).Format("2006-01-02") + FileSuffix

	logier, err := rotatelogs.New(
		logFullPath+"-"+fileSuffix,
		rotatelogs.WithLinkName(linkPath),         // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(int(save)),   // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return logier
}

func getLogName(logFold, level string) (logpath string, linkpath string) {
	hostname, _ := os.Hostname()
	processName := path.Base(os.Args[0])
	user, _ := user.Current()
	logpath, _ = filepath.Abs(path.Join(logFold, fmt.Sprintf("%s.%s.%s.log.%s", processName, hostname, user.Username, strings.ToUpper(level))))
	linkpath = path.Join(logFold, fmt.Sprintf("%s.%s", processName, strings.ToUpper(level)))
	return
}
