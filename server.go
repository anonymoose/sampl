// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Launch an HTTP server listening on the rpc listen port
func launchHttpServer(conf *Config) {
	http.HandleFunc("/status/", httpStatus)

	fullClientAddr := fmt.Sprintf("%s:%s", conf.ClientAddr, conf.ClientPort)
	log.Println("Http client interface listening on ", fullClientAddr)
	log.Fatal(http.ListenAndServe(fullClientAddr, nil))
}

func httpStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: /status \n", r.RemoteAddr)
	fmt.Fprintf(w, serverStatus())
}

func serverStatus() string {
	return "Status: OK"
}
