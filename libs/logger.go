package libs

import (
	"os"
	"sync"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var once sync.Once

// Initialize 로그 생성
func Initialize() {
	once.Do(func() {
		log = logrus.New()

		addSentryHook(log)

		Formatter := new(logrus.TextFormatter)
		Formatter.TimestampFormat = "2006-01-02 15:04:05"
		Formatter.FullTimestamp = true
		log.SetFormatter(Formatter)
	})
}

func addSentryHook(log *logrus.Logger) {
	dsn := os.Getenv("SENTRY_DSN")
	mode := os.Getenv("Mode")
	version := os.Getenv("Version")

	client, err := raven.New(dsn)
	if err != nil {
		log.Fatal(err)
	}

	client.SetEnvironment(mode)
	client.SetRelease(version)
	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err == nil {
		log.Hooks.Add(hook)
	}
}

func Log(level logrus.Level, message string, args ...interface{}) {
	log.Log(level, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Logf(level logrus.Level, message string, args ...interface{}) {
	log.Logf(level, message, args...)
}

func Debugf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

func Infof(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Warnf(message string, args ...interface{}) {
	log.Warnf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	log.Fatalf(message, args...)
}
