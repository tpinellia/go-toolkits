package logger

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger("./var", "app.log", "debug", 7, false)
	Log.Debugln("debug")
	Log.Infoln("info")
	Log.Warningln("warning")
	Log.Errorln("error")
}
