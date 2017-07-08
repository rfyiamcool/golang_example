package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	// log.SetFormatter(log.TextFormatter)
	log.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile("./debug.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetLevel(log.InfoLevel)
}

func pushLogFields(f interface{}) *log.Fields {
	switch f.(type) {
	case A:
		return &log.Fields{
			"role": f.(A).a,
		}
	default:
		return &log.Fields{
		}
	}
}

type A struct {
	a string
}

func main() {

	a := A{"rui"}
	y := pushLogFields(a)
	log.WithFields(*y).Info("A walrus appears")

	y = pushLogFields("xxxx")
	log.WithFields(*y).Info("A walrus appears")
}
