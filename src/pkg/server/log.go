package server

import (
	"log"
	"os"
)

var serverLog *log.Logger

func initLog() {
	serverLog = log.New(os.Stdout,
		"Server: ",
		log.Ldate|log.Ltime)
}

func logf(format string, args ...interface{}) {
	if serverLog == nil {
		initLog()
	}
	serverLog.Printf(format, args...)
}
