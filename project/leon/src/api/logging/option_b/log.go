package optionb

// https://www.sobyte.net/post/2022-03/uber-zap-advanced-usage/ (good read)
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	zapConfig := zap.Config{
		OutputPaths: []string{"stderr"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Log, err = zapConfig.Build()
	if err != nil {
		panic(err)
	}
}

func Debug(msg string, tags ...zap.Field) {

	Log.Debug(msg, tags...)
	Log.Sync()
}
func Info(msg string, tags ...zap.Field) {
	tags = append(tags, Feild("status", "ok")) // default tag
	Log.Info(msg, tags...)
	Log.Sync()
}
func Feild(key string, value any) zap.Field {
	return zap.Any(key, value)
}
