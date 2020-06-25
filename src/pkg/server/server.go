package server

import (
	"net/http"
	"sync"
)

func listen(address string, wg sync.WaitGroup) {
	defer wg.Done()

	err := http.ListenAndServe(address, nil)
	if err != nil && err != http.ErrServerClosed {
		logf(err.Error())
	} else {
		logf("Server closed.")
	}
}

func RunServer(wg *sync.WaitGroup) {
	defer wg.Done()

	logf("Running server...")
	setHandlers()

	address := getServerAddress()
	logf("Address: " + address)

	waitForClose := sync.WaitGroup{}
	waitForClose.Add(1)

	go listen(address, waitForClose)

	logf("Server is listening...")
	waitForClose.Wait()
}
