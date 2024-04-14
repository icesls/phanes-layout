package logger

import (
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func newBufferedWriteSyncer(size int, duration time.Duration, w io.Writer) zapcore.WriteSyncer {
	bufferedSyncer := &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(w),
		Size:          size,
		FlushInterval: duration,
	}
	return bufferedSyncer
}
