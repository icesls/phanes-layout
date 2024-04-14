package logger

import (
	"go.uber.org/zap/zapcore"
	"io"
	"phanes/config"
)

func Init() func() {
	var (
		l       = config.Conf.Collect.Log
		writers = make([]io.Writer, 0, 0)
	)
	if l.FileName == "" {
		panic("no log storage target")
	}

	if l.FileName != "" {
		writers = append(writers, fileOutputWriter("./logs", l.FileName, 500, 3))
	}
	//writers = append(writers, os.Stderr)

	// set your log level here
	logger := newZapLog(zapcore.Level(l.LogLevel), writers...)
	initLogger(logger)

	return func() {
		logger.logger.Sync()
	}
}
