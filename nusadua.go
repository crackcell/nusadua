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
	"github.com/crackcell/kihaadhoo/log"
	"github.com/crackcell/nusadua/config"
	"github.com/crackcell/nusadua/shepherd"
	"sync"
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

	log.Infof("started as %s", config.Role)
	wg.Wait()
}
