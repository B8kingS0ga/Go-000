package main

import (
	"os"
)

func main() {

	e := InitializeServer()
	e.Start()

	signalChan := make(chan os.Signal, 1)

	select {
	case <-signalChan:
		return
	}
}
