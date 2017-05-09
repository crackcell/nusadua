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
 * @file service_discover.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue May  9 14:10:42 2017
 *
 **/

package cluster

import (
	"github.com/crackcell/nusadua/log"
	consulapi "github.com/hashicorp/consul/api"
)

//===================================================================
// Public APIs
//===================================================================

type ServiceDiscover struct {
	consulClient *consulapi.Client
}

func NewServiceDiscover(endpoint string) (clusterInfo *ServiceDiscover, err error) {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = endpoint
	c, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.AppLog.Panicf("consul failed, %s", err)
		return nil, err
	}
	return &ServiceDiscover{consulClient: c}, nil
}

func (this *ServiceDiscover) Register(serviceName string, instanceId string,
	port int) (err error) {

	reg := &consulapi.AgentServiceRegistration{
		Name:    serviceName,
		ID:      instanceId,
		Port:    port,
	}
	return this.consulClient.Agent().ServiceRegister(reg)
}

func (this *ServiceDiscover) DeRegister(instanceId string) error {
	return this.consulClient.Agent().ServiceDeregister(instanceId)
}

//===================================================================
// Private
//===================================================================
