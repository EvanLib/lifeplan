package main

import (
	"github.com/evanlib/lifeplan/srv/lifeplan-api/pkg/api"
	"github.com/evanlib/lifeplan/srv/lifeplan-api/pkg/conf"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/ProtocolONE/rbac"
	redisadapter "github.com/casbin/redis-adapter"
	"github.com/kelseyhightower/envconfig"
	"fmt"
)

func main() {

	
	config := &conf.Config{}
	if err := envconfig.Process("LIFEPLANAPI", config); err != nil {
		log.Fatalf("Config init failed with error: %s\n", err)
	}
	
	// rbac 
	adapter := redisadapter.NewAdapter("tcp", fmt.Sprintf("%s:%d", config.Enforcer.Host, config.Enforcer.Port))
	enf := rbac.NewEnforcer(adapter)
	
	// options for api
	apiOptions := &api.ApiOptions {
		Enforcer: enf,
	}

	// create api server instance
	api, err := api.NewServer(apiOptions)
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