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

package shepherd

import (
	"github.com/crackcell/nusadua/shepherd/rpc"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Rpc struct {
	host string
	port int
	server *thrift.TSimpleServer
	stop chan bool
}

func NewRpc() *Rpc {
	return &Rpc{
		stop: make(chan bool),
	}
}

func (this *Rpc) Start(addr string, port int) (err error) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(fmt.Sprintf("%s:%d", this.host, this.port))
	if err != nil {
		return err
	}

	processorFactory := thrift.NewTProcessorFactory(rpc.NewShepherdServiceProcessor(this))

	this.server = thrift.NewTSimpleServerFactory4(processorFactory, serverTransport, transportFactory, protocolFactory)
	go func() {
		this.server.Serve()
		this.stop <- true
	}()
	return nil
}

func (this *Rpc) Stop() error {
	return this.server.Stop()
}

func (this *Rpc) Wait() {
	<- this.stop
}

func (this *Rpc) SetNodes(nodes []string) (ex *rpc.ShepherdException, err error) {
	return nil, nil
}

func (this *Rpc) GetNodesByFeature(key [][]int64) (r []string, ex *rpc.ShepherdException, err error) {
	return []string{}, nil, nil
}

//===================================================================
// Private
//===================================================================
