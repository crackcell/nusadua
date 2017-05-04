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
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	DataReplicaNum int    `yaml:"data_replica_num"`
}

var Conf = &Config{}

func Parse() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")

	flag.StringVar(&ConfigFile, "r", "follwer", "Node role")
	flag.StringVar(&ConfigFile, "role", "follwer", "Node role")
	flag.StringVar(&ConfigFile, "c", "", "Config file")
	flag.StringVar(&ConfigFile, "config", "", "Config file")

	flag.Parse()
	if Help {
		ShowHelp(0)
	}
	if len(ConfigFile) == 0 {
		ShowHelp(1)
	}

	filename, _ := filepath.Abs(ConfigFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", Conf)
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

    -r, --role         Node role: leader or follower, default: follower
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
