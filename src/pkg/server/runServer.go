package server

import (
	"net/http"
	"sync"
)

func listen(address string, wg sync.WaitGroup) {
	defer wg.Done()

	err := http.ListenAndServe(address, nil)
	if err != nil && err != http.ErrServerClosed {
		logMsg(err.Error())
	} else {
		logMsg("Server closed.")
	}
}

func RunServer(wg *sync.WaitGroup) {
	defer wg.Done()

	logMsg("Running server...")
	setHandlers()

	address := getServerAddress()
	logMsg("Address: " + address)

	waitForClose := sync.WaitGroup{}
	waitForClose.Add(1)

	go listen(address, waitForClose)

	logMsg("Server is listening...")
	waitForClose.Wait()
}
