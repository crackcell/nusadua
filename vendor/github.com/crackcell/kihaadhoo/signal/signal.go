/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Signal handler set
 *
 * @file signal.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 11 15:39:14 2014
 *
 **/

package signal

import (
	"fmt"
	"os"
	"os/signal"
)

//===================================================================
// Public APIs
//===================================================================

type SignalHandler func(s os.Signal, arg interface{})

type SignalHandlerSet interface {
	Start()
	Register(s os.Signal, h SignalHandler)
}

type signalHandlerSet struct {
	m  map[os.Signal]SignalHandler
}

func NewSignalHandlerSet() SignalHandlerSet {
	return &signalHandlerSet{
		m: make(map[os.Signal]SignalHandler),
	}
}

//===================================================================
// Private
//===================================================================

func (this *signalHandlerSet) Start() {
	go func() {
		for {
			c := make(chan os.Signal)
			var sigs []os.Signal
			for sig := range this.m {
				sigs = append(sigs, sig)
			}
			signal.Notify(c, sigs...)
			sig := <-c
			if err := this.handle(sig, nil); err != nil {
				fmt.Printf("unknown signal received: %v, exit unexpectedly\n", sig)
				os.Exit(1)
			}

		}
	}()
}

func (this *signalHandlerSet) Register(s os.Signal, h SignalHandler) {
	this.m[s] = h
}

func (this *signalHandlerSet) handle(s os.Signal, arg interface{}) error {
	if _, found := this.m[s]; found {
		this.m[s](s, arg)
		return nil
	} else {
		return fmt.Errorf("no handler available for signal %v", s)
	}
}
