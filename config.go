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

func chkerr(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
		panic(e)
	}
}

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
	ListenAddr string `yaml:"listen_addr"`
	ListenPort string `yaml:"listen_port"`
}

// Parse the config file specified on the command line into a structure for passing around later.
func parseConfig() *Config {
	cmd := initCommandLine()
	fileContents, e := ioutil.ReadFile(cmd.configFile)
	chkerr(e)

	var conf Config
	err := yaml.Unmarshal(fileContents, &conf)
	chkerr(err)

	return &conf
}
