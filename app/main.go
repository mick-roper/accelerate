package main

import (
	"flag"
	"fmt"
	"log"
)

type iLogger interface {
	Debug(x string)
	Info(x string)
	Fatal(x string)
}

var targetDistance = flag.Float64("distance", -1, "the distance you want to travel")
var acceleration = flag.Float64("acceleration", -1, "the amount of linear acceleration")
var logLevel = flag.String("log-level", "info", "sets the log level")
var withDeceleration = flag.Bool("with-deceleration", false, "(true) to include deceleration in the plan")

var logger iLogger
var calc func(float64, float64, float64, float64) (float64, float64)

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

	if *withDeceleration {
		calc = calculateTransitWithDeceleration
	} else {
		calc = calculateTransit
	}

	var seconds float64 = 0
	var travelled float64 = 0
	var speed float64 = 0

	logger.Info("starting iterations...")

	for travelled < *targetDistance {
		seconds++
		speed, travelled = calc(speed, travelled, *acceleration, *targetDistance)

		logger.Debug(fmt.Sprintf("%fm @ %fms", travelled, speed))
	}

	log.Printf("target reached in %f seconds OR %f minutes OR %f hours", seconds, seconds/60, seconds/60/60)
}

func calculateTransit(speed, travelled, acceleration, targetDistance float64) (float64, float64) {
	newSpeed := speed + acceleration
	return newSpeed, travelled + newSpeed
}

func calculateTransitWithDeceleration(speed, travelled, acceleration, targetDistance float64) (float64, float64) {
	var newSpeed float64
	var newAcc float64
	
	if acceleration > targetDistance - travelled {
		newAcc = (targetDistance - travelled) / 2
	} else {
		newAcc = acceleration
	}

	if travelled >= targetDistance*0.50 {
		newAcc = newAcc * -1
	}

	newSpeed = speed + newAcc

	if newSpeed < 0 {
		newSpeed = 1
	}

	return newSpeed, travelled + newSpeed
}

type debugLogger struct{}

func (l *debugLogger) Debug(x string) {
	log.Print(x)
}

func (l *debugLogger) Info(x string) {
	log.Print(x)
}

func (l *debugLogger) Fatal(x string) {
	log.Fatal(x)
}

type infoLogger struct{}

func (l *infoLogger) Debug(x string) {}

func (l *infoLogger) Info(x string) {
	log.Print(x)
}

func (l *infoLogger) Fatal(x string) {
	log.Fatal(x)
}
