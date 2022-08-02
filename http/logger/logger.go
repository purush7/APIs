package logger

import log "github.com/sirupsen/logrus"

type LogInstance interface {
	Error() string
}

func LogError(name string, obj LogInstance) {
	if obj == nil {
		return
	}
	log.Printf("[%s]: %s\n", name, obj.Error())
}

func FatalError(name string, obj LogInstance) {
	log.Fatalf("[%s]: %s\n", name, obj.Error())
}

func Debug(name string, obj interface{}) {
	log.Debug("name", obj)
}
