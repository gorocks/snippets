package main

import (
	"github.com/sirupsen/logrus"
)

const timeFmt = "2006/01/02 15:04:05"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeFmt,
	})
	logrus.Info("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeFmt,
	})
	logrus.WithField("foo", "bar").Info()
}
