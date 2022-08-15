package optiona

import (
	"example/config"
	"fmt"
	"os"
	"strings"

	logrus "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	fmt.Println("Logging")
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
	//Log.SetFormatter(&logrus.TextFormatter{})
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseTags(tags...)).Info(msg)
}
func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseTags(tags...)).Debug(msg)
}

func parseTags(tags ...string) logrus.Fields {
	results := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		key, value := strings.Split(tag, ":")[0], strings.Split(tag, ":")[1]
		results[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	return results
}
