package main

import (
	"example.com/tracker/internal/app"
	"github.com/sirupsen/logrus"
)

// @title Time Tracker
// @version 1.0
// @description API Server for Time Tracker

// @host localhost:3000
// BasePath /
func main() {
	err := app.SetupConfig("config/.env")
	if err != nil {
		logrus.Fatal(err)
	}
	err = app.SetupLogger()
	if err != nil {
		logrus.Fatal(err)
	}
	err = app.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
