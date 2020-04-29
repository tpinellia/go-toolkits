package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//LogFormatter 日志自定义格式
type LogFormatter struct{}

//Format 格式化日志信息
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var msg string
	if entry.HasCaller() {
		msg = fmt.Sprintf(
			"%s [%s:%d][%s] %s\n",
			entry.Time.Local().Format("2006-01-02 15:04:05"),
			filepath.Base(entry.Caller.File),
			entry.Caller.Line,
			strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
	}
	msg = fmt.Sprintf("%s [%s] %s\n", entry.Time.Local().Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

//New 初始化logger
func New(logPath string, logFile string, level string, logAge int, caller bool) (*logrus.Logger, error) {
	var Log = logrus.New()
	switch level {
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "warning":
		Log.SetLevel(logrus.WarnLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "trace":
		Log.SetLevel(logrus.TraceLevel)
	default:
		Log.SetLevel(logrus.DebugLevel)
		Log.Warningln("log level only allow [panic,fatal,error,warning,info,debug,trace].your choice is wrong,so use the default level 'debug'.")
	}

	if logAge < 7 {
		logAge = 7
		Log.Warningln("use the recommended logAge 7.")
	}

	if caller {
		Log.SetReportCaller(true)
	}

	_, err := os.Stat(logPath)
	if !(err == nil || os.IsExist(err)) {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			return Log, errors.Errorf("mkdir dir %s failed,%v", logPath, err)
		}
	}
	if err := syscall.Access(logPath, syscall.O_RDWR); err != nil {
		return Log, errors.Errorf("access dir %s error,%v", logPath, err)
	}
	baseLog := path.Join(logPath, logFile)
	writer, err := rotatelogs.New(
		baseLog+".%Y%m%d",
		rotatelogs.WithLinkName(baseLog), //生成软链，指向最新的日志文件
		rotatelogs.WithMaxAge(time.Duration(logAge*24*60*60)*time.Second), // 默认为7天，文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(24*60*60)*time.Second),  // 默认为1天，日志切割时间间隔
	)
	if err != nil {
		return Log, errors.Errorf("config logger error,%v", err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, new(LogFormatter))
	Log.AddHook(lfHook)
	return Log, nil
}
