package logger

import "log"

type LogInstance interface {
	Error() string
}

func Logger(name string, obj LogInstance) {
	if obj == nil {
		return
	}
	log.Printf("[%s]: %s\n", name, obj.Error())
}

func LoggerAndKiller(name string, obj LogInstance) {
	log.Fatalf("[%s]: %s\n", name, obj.Error())
}
