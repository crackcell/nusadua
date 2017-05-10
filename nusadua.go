/***************************************************************
 *
 * Copyright (c) 2017, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the Apache licence
 *
 **************************************************************/

/**
 *
 *
 * @file nusadua.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed May  3 15:31:46 2017
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/kihaadhoo/signal"
	"github.com/crackcell/nusadua/cluster"
	"github.com/crackcell/nusadua/config"
	"github.com/crackcell/nusadua/log"
	"github.com/crackcell/nusadua/router"
	"github.com/crackcell/nusadua/server"
	"os"
	"sync"
	"syscall"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

var wg sync.WaitGroup
var routerRpc = router.NewRpc()
var serverRpc = server.NewRpc()
var discover *cluster.Catalog
var instanceId string

func runRpc() {
	if config.Role == "router" {
		routerRpc.Start(config.GlobalConfig.RouterConfig.Host,
			config.GlobalConfig.RouterConfig.Port)

		wg.Add(1)
		go func() {
			defer wg.Done()
			routerRpc.Wait()
		}()
	}

	if config.Role == "server" {
		serverRpc.Start(config.GlobalConfig.ServerConfig.Host,
			config.GlobalConfig.ServerConfig.Port)

		wg.Add(1)
		go func() {
			defer wg.Done()
			serverRpc.Wait()
		}()
	}
}

func registerService() {
	// register to cluster discover service
	var err error
	discover, err = cluster.NewServiceDiscover(config.GlobalConfig.ConsulConfig.AgentAddr)
	if err != nil {
		panic(err)
	}
	instanceId = fmt.Sprintf("%s:%d",
		config.GlobalConfig.ConsulConfig.ServiceName,
		config.GlobalConfig.ServerConfig.Port)
	discover.Register(config.GlobalConfig.ConsulConfig.ServiceName,
		instanceId,
		config.GlobalConfig.ServerConfig.Port)
}

func handleSignals(cleanup func()) {
	// init signal handlers
	sset := signal.NewSignalHandlerSet()

	handler := func(s os.Signal, arg interface{}) {
		log.AppLog.Infof("received signal: %v", s)
		if s == syscall.SIGTERM {
			log.AppLog.Infof("signal terminate received, exited normally")
			cleanup()
		}
	}
	sset.Register(syscall.SIGINT, handler)
	sset.Register(syscall.SIGUSR1, handler)
	sset.Register(syscall.SIGUSR2, handler)
	sset.Register(syscall.SIGTERM, handler)

	sset.Start()
}

func main() {
	config.Init()
	log.Init()

	runRpc()
	registerService()

	cleanup := func() {
		if config.Role == "router" {
			routerRpc.Stop()
		}
		if config.Role == "server" {
			serverRpc.Stop()
		}
		discover.DeRegister(instanceId)
	}
	handleSignals(cleanup)

	// wait for finish
	log.AppLog.Infof("started as %s", config.Role)
	wg.Wait()
}
