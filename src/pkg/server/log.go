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

func logMsg(msg string) {
	if serverLog == nil {
		initLog()
	}
	serverLog.Println(msg)
}
