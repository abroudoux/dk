package logs

import "github.com/charmbracelet/log"

func InfoMsg(msg string) {
	log.Info(msg)
}

func WarnMsg(msg string) {
	log.Warn(msg)
}

func Warn(msg string, err error) {
	log.Warn(msg, err)
}

func ErrorMsg(msg string) {
	log.Error(msg)
}

func Error(msg string, err error) {
	log.Error(msg, err)
}

func Fatal(msg string, err error) {
	log.Fatal(msg, err)
}