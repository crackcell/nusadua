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
 * @file rpc.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri May  5 17:40:17 2017
 *
 **/

package router

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/crackcell/nusadua/config"
	"github.com/crackcell/nusadua/router/rpc"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type Rpc struct {
	lock         *sync.Mutex
	started      bool
	server       *thrift.TSimpleServer
	stop         chan bool
	featureShard *FeatureShard
}

func NewRpc() *Rpc {
	return &Rpc{
		lock:         new(sync.Mutex),
		started:      false,
		stop:         make(chan bool),
		featureShard: NewFeatureShard([]string{}, 3),
	}
}

func (this *Rpc) Start(host string, port int) (err error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	addr := fmt.Sprintf("%s:%d", host, port)
	serverTransport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	processorFactory := thrift.NewTProcessorFactory(rpc.NewRouterServiceProcessor(this))

	this.server = thrift.NewTSimpleServerFactory4(processorFactory, serverTransport, transportFactory, protocolFactory)
	go func() {
		this.server.Serve()
		this.stop <- true
	}()

	this.started = true
	return nil
}

func (this *Rpc) Stop() (err error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if !this.started {
		return nil
	}
	if err = this.server.Stop(); err != nil {
		return err
	}
	this.started = false
	return nil
}

func (this *Rpc) Wait() {
	if !this.started {
		return
	}
	<-this.stop
}

func (this *Rpc) SetNodes(nodes []string) (ex *rpc.RouterException, err error) {
	this.featureShard.SetNodes(nodes, config.GlobalConfig.RouterConfig.ReplicaNum)
	return nil, nil
}

func (this *Rpc) GetNodesByFeature(key []int64) (r []string, ex *rpc.RouterException, err error) {
	if nodes, err := this.featureShard.GetNodesByFeature(key); err != nil {
		ex := rpc.NewRouterException()
		ex.Message = err.Error()
		ex.Status = 1
		return nodes, ex, err
	} else {
		return nodes, nil, nil
	}
}

//===================================================================
// Private
//===================================================================
