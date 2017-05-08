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
 * @file log.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon May  8 15:39:03 2017
 *
 **/

package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/crackcell/nusadua/config"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

var AccessLog = logrus.New()
var AppLog = logrus.New()
var StatLog = logrus.New()

func Init() (err error) {
	// init access log
	AccessLog.Formatter = &logrus.JSONFormatter{}
	AccessLog.Level, err = logrus.ParseLevel(config.GlobalConfig.LogConfig.AccessLogLevel)
	if err != nil {
		AccessLog.Level = logrus.InfoLevel
	}
	AccessLog.Out = os.Stdout

	// init app log
	AppLog.Formatter = &logrus.JSONFormatter{}
	AppLog.Level, err = logrus.ParseLevel(config.GlobalConfig.LogConfig.AppLogLevel)
	if err != nil {
		AppLog.Level = logrus.InfoLevel
	}
	AppLog.Out = os.Stdout

	// init stat log
	StatLog.Formatter = &logrus.JSONFormatter{}
	StatLog.Level, err = logrus.ParseLevel(config.GlobalConfig.LogConfig.StatLogLevel)
	if err != nil {
		StatLog.Level = logrus.InfoLevel
	}
	StatLog.Out = os.Stdout

	return nil
}

//===================================================================
// Private
//===================================================================
