package server

import (
	"net/http"
	"pkg/database"
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

func RunServer(dbUrl string, wg *sync.WaitGroup) {
	defer func() {
		logf("Disposing database...")
		if err := database.DisposeDataBase(); err != nil {
			logf("Cannot dispose database: %s.", err.Error())
		}
		wg.Done()
	}()

	if !validateDbUrl(dbUrl) {
		logf("Invalid dbUrl provided.")
		return
	}

	logf("Running server...")
	setHandlers()

	logf("Connecting do database...")
	err := database.InitDataBase(dbUrl)
	if err != nil {
		logf("Cannot reach database: %s.", err.Error())
		return
	}
	logf("Database is ready.")

	address := getServerAddress()
	logf("Address: " + address)

	waitForClose := sync.WaitGroup{}
	waitForClose.Add(1)

	go listen(address, waitForClose)

	logf("Server is listening...")
	waitForClose.Wait()
}
