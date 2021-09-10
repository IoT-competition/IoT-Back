package basic

import (
	"github.com/google/wire"
	"os"

	"github.com/IoTCompetition/IoBack/app/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LoggerSet logger DI
var LoggerSet = wire.NewSet(GetLogWriter, GetEncoder, NewLogger)

/*
NewLogger 获得 logger
*/
func NewLogger(c *config.Type, writeSync zapcore.WriteSyncer, encoder zapcore.Encoder) *zap.Logger {
	atomicLevel := zap.NewAtomicLevel()
	if c.GetAppConfig().GetMode() == "dev" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSync)),
		atomicLevel,
	)

	var logger *zap.Logger
	if c.GetAppConfig().GetMode() == "dev" {
		logger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		logger = zap.New(core)
	}

	return logger
}

// GetLogWriter 日志切割
func GetLogWriter(c *config.Type) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   c.GetLogConfig().GetFilename(),
		MaxSize:    c.GetLogConfig().GetMaxSize(),
		MaxBackups: c.GetLogConfig().GetMaxBackups(),
		MaxAge:     c.GetLogConfig().GetMaxAge(),
		Compress:   c.GetLogConfig().GetCompress(),
	}

	return zapcore.AddSync(hook)
}

// GetEncoder 进行个性化编码
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
