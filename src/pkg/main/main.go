package main

import (
	"fmt"
	"pkg/server"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go server.RunServer(wg)
	wg.Wait()

	fmt.Println("Termination...")
}
