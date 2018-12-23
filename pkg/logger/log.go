package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gmdb/pkg/config"
)

var log *logrus.Logger

func Initialize() {
	//filePath := getLogFilePath()
	//fileName := getLogFileName()

	log = logrus.New()
}

func Trace(s string) {
	log.Trace(s)
}

func Debug(s string) {
	log.Debug(s)
}

func Info(s string) {
	log.Info(s)
}

func Warn(s string) {
	log.Warn(s)
}

func Error(s string) {
	log.Error(s)
}

// Calls os.Exit(1) after logging
func Fatal(s string) {
	log.Fatal(s)
}

// Calls panic() after logging
func Panic(s string) {
	log.Panic(s)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", "test/", config.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.App.LogSaveName,
		time.Now().Format(config.App.TimeFormat),
		config.App.LogFileExt,
	)
}
