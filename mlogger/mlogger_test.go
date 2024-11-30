package mlogger

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger()
	MLogger.Info("Hello World")
}
