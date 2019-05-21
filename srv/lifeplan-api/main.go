package main

import (
	"github.com/evanlib/lifeplan/srv/lifeplan-api/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create api server instance
	api, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	//initialize server
	api.Start()


	handleOsSignals(api)

}

func handleOsSignals(server *api.Api) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	exitChan := make(chan int)

	go func() {
		for {
			s := <-signalChan
			switch s {
			case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
				server.Stop()
				exitChan <- 0
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}