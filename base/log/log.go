package log

import (
	"github.com/ijkbytes/mega/base/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

func toLevel(lv int) zap.AtomicLevel {
	switch lv {
	case 0:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case 1:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case 2:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case 3:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case 4:
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case 5:
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case 6:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	}
}

func Init() *zap.Logger {
	cores := []zapcore.Core{}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.StacktraceKey = "stack"
	encoderCfg.TimeKey = "@timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewJSONEncoder(encoderCfg)

	fileOutPutter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Mega.Log.Path,
		MaxSize:    config.Mega.Log.MaxSize,
		MaxBackups: config.Mega.Log.MaxBackups,
		MaxAge:     config.Mega.Log.MaxAge,
		LocalTime:  true,
		Compress:   true,
	})

	lv := toLevel(config.Mega.Log.Level)
	cores = append(cores, zapcore.NewCore(enc, fileOutPutter, lv))

	// console
	if config.Mega.Log.Console {
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})

		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleErrors, highPriority))
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	}

	core := zapcore.NewTee(cores...)

	opts := []zap.Option{}
	if config.Mega.Log.StackTrace {
		opts = append(opts, zap.AddStacktrace(lv))
	}

	logger = zap.New(core, opts...)

	return logger
}

func Get(name string) *zap.Logger {
	return logger.Named(name)
}
