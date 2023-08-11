package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	// 红色
	errLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	// 蓝色
	infoLog = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errLog, infoLog}
	mu      sync.Mutex
)

var (
	Error  = errLog.Println
	Errorf = errLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ErrorLevel < level {
		errLog.SetOutput(io.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}

func main() {
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
