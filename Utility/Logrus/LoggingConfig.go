package Logrus

import (
	"fmt"
	"github.com/kz/discordrus"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)



func init() {

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: time.ANSIC,
	})

	logger.SetOutput(os.Stderr)

	logger.AddHook(discordrus.NewHook(

		os.Getenv("DISCORDRUS_WEBHOOK_URL"),
		logrus.WarnLevel,
		&discordrus.Opts{
			Username:           "Logrus",
			EnableCustomColors: true,
			CustomLevelColors: &discordrus.LevelColors{
				Debug: 10170623,
				Info:  3581519,
				Warn:  14327864,
				Error: 13631488,
				Panic: 13631488,
				Fatal: 13631488,
			},
			DisableInlineFields: false,
		},
	))


}



var logger = logrus.New()

func WithFields(fields logrus.Fields) *logrus.Entry{
	return logger.WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry{
	return logger.WithField(key, value)
}

func WithLocation() *logrus.Entry{

	fpcs := make([]uintptr, 1)
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		fmt.Println("MSG: NO CALLER")
	}

	caller := runtime.FuncForPC(fpcs[0]-1)
	if caller == nil {
		fmt.Println("MSG CALLER WAS NIL")
	}

	file, line := caller.FileLine(fpcs[0]-1)
	return logger.WithField("Function", caller.Name()).WithField("File", file).WithField("Line", line)
}

func Info(args ...interface{}){
	logger.Info(args)
}

func Error(args ...interface{}){
	logger.Error(args)
}

func Debug(args ...interface{}){
	logger.Debug(args)
}

func Warn(args ...interface{}){
	logger.Warn(args)
}

func Fatal(args ...interface{}){
	logger.Fatal(args)
}
