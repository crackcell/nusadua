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
 * @date Tue May  9 11:57:24 2017
 *
 **/

package server

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/crackcell/nusadua/rpc"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type Rpc struct {
	lock    *sync.Mutex
	started bool
	server  *thrift.TSimpleServer
	stop    chan bool
}

func NewRpc() *Rpc {
	return &Rpc{
		lock:    new(sync.Mutex),
		started: false,
		stop:    make(chan bool),
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

	processorFactory := thrift.NewTProcessorFactory(rpc.NewParameterServerProcessor(this))

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

func (this *Rpc) CreateTask(name string, maxKey int64) (ex *rpc.RpcException, err error) {
	return nil, nil
}

func (this *Rpc) MultiPush(keys [][]int64, values []float64) (ex *rpc.RpcException, err error) {
	return nil, nil
}

func (this *Rpc) MultiPull(keys [][]int64) (r []float64, ex *rpc.RpcException, err error) {
	return []float64{}, nil, nil
}

func (this *Rpc) RangePush(start_key int64, end_key int64, values []float64) (ex *rpc.RpcException, err error) {
	return nil, nil
}

func (this *Rpc) RangePull(start_key int64, end_key int64) (r []float64, ex *rpc.RpcException, err error) {
	return []float64{}, nil, nil
}

//===================================================================
// Private
//===================================================================
