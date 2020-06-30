package main

import (
	"fmt"
	"os"
	"pkg/server"
	"sync"
)

const (
	ErrCode = -1
	usage   = "Usage:\nserver.exe <db_connection_string>"
)

func showUsage() {
	fmt.Println(usage)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of arguments.")
		showUsage()
		os.Exit(ErrCode)
	}

	dbUrl := args[0]

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go server.RunServer(dbUrl, wg)
	wg.Wait()

	fmt.Println("Termination...")
}
