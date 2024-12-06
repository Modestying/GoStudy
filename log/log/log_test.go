package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(InfoLevel)

	Error("sas")
	Errorf("%sn", "error")

	Info("info")
	Infof("%s\n", "info")

	SetLevel(ErrorLevel)
	Error("sas")
	Errorf("%sn", "error")

	Info("info")
	Infof("%s\n", "info")
}
