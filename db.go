// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"
)

//
// Launch a pool of workers so we can queue work to them.
//  http://golang.org/doc/codewalk/sharemem/
//
func launchDbHandlerPool(conf *Config) {
	log.Printf("Launching DB handler pool\n")
}

//
// Make sure everything is ok to proceed.
// 1. Ensure that data directory is present and writable.
//
func ensureEnvironment(conf *Config) {
	log.Printf("Checking Environment...\n")

	ensureDataDir(conf)

	log.Printf("done.\n")
}

//
// check for presence of config::data_dir.  If it isn't there, create it.
//
func ensureDataDir(conf *Config) {
	if _, err := os.Stat(conf.DataDir); os.IsNotExist(err) {
		log.Printf("    config::data_dir '%s' does not exist.  Attempting to create...", conf.DataDir)
		err := os.MkdirAll(conf.DataDir, 0750)
		chkerr(err)
	} else {
		log.Printf("    config::data_dir '%s' exists.", conf.DataDir)
	}
}

const COMMAND_WRITE = 0
const COMMAND_READ = 1
const COMMAND_DELETE = 2

//
// Contents that we are trying to write to disk.
//
type Payload struct {
	content *map[string][]string
}

//
// Request from the client.  Can be one of the const's above.
//
type Command struct {
	commandType  int
	database     string
	collection   string
	instructions *map[string]string
	payload      *Payload
	timestamp    *time.Time
}

func makeCommand(database string, collection string, params *map[string]string, payloadContents *map[string][]string) *Command {
	now := time.Now().UTC()
	return &Command{
		commandType:  COMMAND_WRITE,
		database:     database,
		collection:   collection,
		instructions: params,
		payload:      &Payload{content: payloadContents},
		timestamp:    &now,
	}
}

//
// Return information on general status.
//
func dbStatus() string {
	return "Status: OK"
}

//
// Take an write instruction and save it to file.
//
func dbWrite(database string, collection string, params *map[string]string, payloadContents *map[string][]string) string {
	command := makeCommand(database, collection, params, payloadContents)

	return "Write: OK"
}
