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
 * @file config.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed May  3 15:25:34 2017
 *
 **/

package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//===================================================================
// Public APIs
//===================================================================

var (
	Help    bool
	Verbose bool

	Role string

	ConfigFile string
)

type Config struct {
	LogConfig      LogConfig      `yaml:"log"`
	ShepherdConfig ShepherdConfig `yaml:"shepherd"`
}

type LogConfig struct {
	AccessLogPath  string `yaml:"access_log_path"`
	AccessLogLevel string `yaml:"access_log_level"`
	AppLogPath     string `yaml:"app_log_path"`
	AppLogLevel    string `yaml:"app_log_level"`
	StatLogPath    string `yaml:"stat_log_path"`
	StatLogLevel   string `yaml:"stat_log_level"`
}

type ShepherdConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	DataReplicaNum int    `yaml:"data_replica_num"`
}

var GlobalConfig = &Config{}

func Init() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")

	flag.StringVar(&Role, "r", "server", "Node role")
	flag.StringVar(&Role, "role", "server", "Node role")
	flag.StringVar(&ConfigFile, "c", "", "Config file")
	flag.StringVar(&ConfigFile, "config", "", "Config file")

	flag.Parse()
	if Help {
		ShowHelp(0)
	}
	if len(ConfigFile) == 0 {
		fmt.Println("wrong argument: --config")
		ShowHelp(1)
	}

	if len(Role) == 0 && Role != "shepherd" && Role != "server" {
		fmt.Println("wrong argument: --role")
		ShowHelp(1)
	}

	filename, _ := filepath.Abs(ConfigFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		panic(err)
	}
}

//===================================================================
// Private
//===================================================================

const (
	logoString = ` _______                     _____              
|    |  |.--.--.-----.---.-.|     \.--.--.---.-.
|       ||  |  |__ --|  _  ||  --  |  |  |  _  |
|__|____||_____|_____|___._||_____/|_____|___._|
`
	helpString = `Run parameter server
Usage:
    nusadua [options]
Options:
    -h, --help         Print this message
    -v, --verbose      Use verbose output

    -r, --role         Node role: shepherd or server, default: server
	-c, --config       Config file path
`
)

func ShowHelp(ret int) {
	fmt.Println(logoString)
	fmt.Print(helpString)
	os.Exit(ret)
}

func ShowLogo() string {
	return logoString
}
