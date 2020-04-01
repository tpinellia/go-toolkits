package logger

import (
	"os"
	"path"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//Log 全局log
var Log *logrus.Logger

//InitLogger 初始化logger
func InitLogger(logPath string, logFile string, level string, logAge int, nocolor bool) {
	Log = logrus.New()
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

	_, err := os.Stat(logPath)
	if !(err == nil || os.IsExist(err)) {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			Log.Errorf("mkdir logPath %s failed.%v.", logPath, err)
		}
	}
	baseLog := path.Join(logPath, logFile)
	writer, err := rotatelogs.New(
		baseLog+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Duration(logAge*24*60*60)*time.Second), // 默认为7天，文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(24*60*60)*time.Second),  // 默认为1天，日志切割时间间隔
	)
	if err != nil {
		Log.Errorln("config logger error,", err)

	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &nested.Formatter{TimestampFormat: "2006-01-02 15:04:05", NoColors: nocolor})
	Log.AddHook(lfHook)
}
