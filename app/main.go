package main

import (
	"flag"
	"log"
	"fmt"
)

type iLogger interface {
	Debug(x string)
	Info(x string)
	Fatal(x string)
}

var targetDistance = flag.Float64("distance", -1, "the distance you want to travel")
var acceleration = flag.Float64("acceleration", -1, "the amount of linear acceleration")
var logLevel = flag.String("log-level", "info", "sets the log level")

var logger iLogger

func main() {
	flag.Parse()

	if *logLevel == "debug" {
		logger = &debugLogger{}
	} else {
		logger = &infoLogger{}
	}

	if *targetDistance < 0 {
		log.Fatal("target is invalid")
	}

	if *acceleration < 0 {
		log.Fatal("acceleration is invalid")
	}

	var seconds float64 = 0
	var travelled float64 = 0
	var speed float64 = 0

	logger.Info("starting iterations...")

	for travelled < *targetDistance {
		seconds++
		speed += *acceleration
		travelled += speed

		logger.Debug(fmt.Sprintf("%vm @ %vms", travelled, speed))
	}

	log.Printf("target reached in %f seconds OR %f minutes OR %f hours", seconds, seconds / 60, seconds /60 / 60)
}

type debugLogger struct {}

type infoLogger struct {}

func (l *debugLogger) Debug(x string) {
	log.Print(x)
}

func (l *debugLogger) Info(x string) {
	log.Print(x)
}

func (l *debugLogger) Fatal(x string) {
	log.Fatal(x)
}

func (l *infoLogger) Debug(x string) {}

func (l *infoLogger) Info(x string) {
	log.Print(x)
}

func (l *infoLogger) Fatal(x string) {
	log.Fatal(x)
}