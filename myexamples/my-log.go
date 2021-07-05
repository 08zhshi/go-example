package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

var url string = "http://www.baidu.com"

// In contexts where performance is nice, but not critical, use the SugaredLogger.
// It's 4-10x faster than other structured logging packages and includes both structured and printf-style APIs.
func zapSugarLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// When performance and type safety are critical, use the Logger.
// It's even faster than the SugaredLogger and allocates far less, but it only supports structured logging.
func zapLog() {
	//logger, _ := zap.NewProduction()
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := config.Build()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func main() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	zapSugarLog()
	zapLog()
}
