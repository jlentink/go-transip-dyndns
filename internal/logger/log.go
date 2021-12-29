package logger

import (
	"bytes"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

var _log *logrus.Logger

type outputSplitter struct{}

func (splitter *outputSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}

// Init Logger setup
func init() {
	_log = &logrus.Logger{
		Out:   &outputSplitter{},
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			LogFormat: "%lvl% - %msg%",
		},
	}
}

// Get Logger setup
func Get() *logrus.Logger {
	return _log
}

// SetVerbose level on/off
func SetVerbose(v bool) {
	if v {
		_log.SetLevel(logrus.DebugLevel)
		_log.Debug("Log level set to debug...\n")
	}
}
