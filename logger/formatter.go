package logger

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

const (
	FORMAT_TEXT   = "text"
	FORMAT_JSON   = "json"
	FORMAT_CUSTOM = "custom"
)

const (
	blue       = 36
	bright_red = 41
	pink       = 35
	green      = 32
	red        = 31
	yellow     = 33
)

type GameLogType uint32

const TYPE_UNKNOW GameLogType = 0
const (
	TYPE_REDIS GameLogType = 1 << iota
	TYPE_TCP
	TYPE_ROUTER
	TYPE_CGI
	TYPE_GAME
)

func ParseGameLogType(s string) GameLogType {
	switch strings.ToLower(s) {
	case "redis":
		return TYPE_REDIS
	case "tcp":
		return TYPE_TCP
	case "router":
		return TYPE_ROUTER
	case "cgi":
		return TYPE_CGI
	case "game":
		return TYPE_GAME
	default:
		return TYPE_UNKNOW
	}
}

type GameFormatter struct {
	AppName      string
	showLogLevel uint32
	ForceColors  bool
	isTerminal   bool
}

func NewGameFormater(appName string, types GameLogType) *GameFormatter {
	return &GameFormatter{
		AppName:      appName,
		showLogLevel: uint32(types),
		isTerminal:   true,
	}
}

func (m *GameFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	if entry.Time.IsZero() {
		entry.Time = time.Now()
	}

	var levelColor int
	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = green
	case logrus.DebugLevel:
		levelColor = pink
	case logrus.TraceLevel:
		levelColor = blue
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel:
		levelColor = red
	case logrus.PanicLevel, logrus.FatalLevel:
		levelColor = bright_red
	default:
		levelColor = blue
	}
	level := strings.ToUpper(entry.Level.String())
	level = level[0:4]

	caller := ""
	if entry.HasCaller() {
		_, file := filepath.Split(entry.Caller.File)
		caller = fmt.Sprintf("%s:%d ", file, entry.Caller.Line)
	}

	traceIDString := ""
	if entry.Context != nil {
		traceIDString = fmt.Sprintf("[%s]", trace.SpanContextFromContext(entry.Context).TraceID().String())
	}
	if traceID, ok := entry.Data["traceID"]; ok {
		traceIDString = fmt.Sprintf("[%s]", traceID)
	}
	fmt.Fprintf(b, "[%s] [%s] [\x1b[%dm%s\x1b[0m]%s", m.AppName, entry.Time.Format("2006-01-02 15:04:05.000000"), levelColor, level, traceIDString)
	// fmt.Fprintf(b, "\x1b[%dm%s| %s | %s\x1b[0m", levelColor, levelText, timeFormat, caller)
	for k, v := range entry.Data {
		if k == "traceID" {
			continue
		}
		b.WriteString(fmt.Sprintf(" \x1b[%dm%s\x1b[0m:%v", levelColor, k, v))
	}
	entry.Message = strings.TrimSuffix(entry.Message, "\n")
	b.WriteString(" => ")
	b.WriteString(caller)
	b.WriteString(entry.Message)
	b.WriteString("\n")
	return b.Bytes(), nil
}

func (m *GameFormatter) AddType(t GameLogType) {
	m.showLogLevel = m.showLogLevel | uint32(t)
}
func (m *GameFormatter) RemoveType(t GameLogType) {
	m.showLogLevel = m.showLogLevel &^ uint32(t)
}
