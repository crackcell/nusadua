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
	"os"
	"sync"
	"syscall"

	"github.com/crackcell/kihaadhoo/signal"
	"github.com/crackcell/nusadua/config"
	"github.com/crackcell/nusadua/log"
	"github.com/crackcell/nusadua/shepherd"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

func main() {
	config.Init()
	log.Init()

	var wg sync.WaitGroup

	sheperdRpc := shepherd.NewRpc()
	sheperdRpc.Start(config.GlobalConfig.ShepherdConfig.Host, config.GlobalConfig.ShepherdConfig.Port)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sheperdRpc.Wait()
	}()

	// init signal handlers
	sset := signal.NewSignalHandlerSet()

	cleanup := func(s os.Signal, arg interface{}) {
		log.AppLog.Infof("received signal: %v", s)
		if s == syscall.SIGTERM {
			log.AppLog.Infof("signal terminate received, exited normally")
			sheperdRpc.Stop()
		}
	}
	sset.Register(syscall.SIGINT, cleanup)
	sset.Register(syscall.SIGUSR1, cleanup)
	sset.Register(syscall.SIGUSR2, cleanup)
	sset.Register(syscall.SIGTERM, cleanup)

	sset.Start()

	// wait for finish
	log.AppLog.Infof("started as %s", config.Role)
	wg.Wait()
}
