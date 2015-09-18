// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const DEFAULT_CONFIG_FILE = "./sampl.conf"

const DEFAULT_LISTEN_ADDR = "127.0.0.1"
const DEFAULT_LISTEN_PORT = "5151"
const DEFAULT_CLIENT_ADDR = "127.0.0.1"
const DEFAULT_CLIENT_PORT = "5152"
const DEFAULT_DATA_DIR = "./data"

// Command line argument parsing.
type CmdLine struct {
	configFile string
}

func initCommandLine() *CmdLine {
	cmd := &CmdLine{configFile: ""}
	flag.StringVar(&cmd.configFile, "config", "./sampl.conf", "Readable path to the config file.")
	flag.Parse()
	return cmd
}

// Runtime configuration
type Config struct {
	ClientAddr string `yaml:"client_addr"`
	ClientPort string `yaml:"client_port"`

	ListenAddr string `yaml:"listen_addr"`
	ListenPort string `yaml:"listen_port"`

	DataDir string `yaml:"data_dir"`
}

//
// Parse the config file specified on the command line into a structure for passing around later.
//
func parseConfig() *Config {
	cmd := initCommandLine()
	fileContents, e := ioutil.ReadFile(cmd.configFile)
	chkerr(e)

	var conf Config
	err := yaml.Unmarshal(fileContents, &conf)
	chkerr(err)

	checkConfig(&conf)

	return &conf
}

//
// Ensure that rules are followed for configuration files.
//
func checkConfig(conf *Config) {
	if conf.ClientPort == conf.ListenPort {
		log.Fatalf("CONFIG FILE ERROR: client_port can not be the same as listen_port")
	}
}
