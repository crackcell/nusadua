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
 * @file catalog.go
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

type Instance struct {
	ID   string
	Host string
	Port int
}

type Catalog struct {
	consulClient *consulapi.Client
}

func NewServiceDiscover(addr string) (clusterInfo *Catalog, err error) {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = addr
	c, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.AppLog.Panicf("consul failed, %s", err)
		return nil, err
	}
	return &Catalog{consulClient: c}, nil
}

func (this *Catalog) Register(serviceName string, instanceId string,
	port int) (err error) {

	reg := &consulapi.AgentServiceRegistration{
		Name: serviceName,
		ID:   instanceId,
		Port: port,
	}
	return this.consulClient.Agent().ServiceRegister(reg)
}

func (this *Catalog) DeRegister(instanceId string) error {
	return this.consulClient.Agent().ServiceDeregister(instanceId)
}

func (this *Catalog) GetInstances(serviceName string) (instances []*Instance, err error) {
	services, _, err := this.consulClient.Catalog().Service(serviceName, "", nil)
	if err != nil {
		return []*Instance{}, nil
	}
	for _, s := range services {
		instances = append(instances, &Instance{ID: s.ServiceID, Host: s.Node, Port: s.ServicePort})
	}
	return instances, nil
}

//===================================================================
// Private
//===================================================================
