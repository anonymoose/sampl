// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	//"net/url"
	"strings"
)

//
// Launch an HTTP server listening on the rpc listen port
//
func launchHttpServer(conf *Config) {
	http.HandleFunc("/status/", httpStatus)
	http.HandleFunc("/write/", httpWrite)

	fullClientAddr := fmt.Sprintf("%s:%s", conf.ClientAddr, conf.ClientPort)
	log.Println("Http client interface listening on ", fullClientAddr)
	log.Fatal(http.ListenAndServe(fullClientAddr, nil))
}

//
// Wrapper around HTTP call to get server status.
//
func httpStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: /status \n", r.RemoteAddr)
	var path string
	var params *[]string
	var data *map[string][]string
	path, params, data = httpParseRequest(r)
	out := httpStatusImpl(path, params, data)
	fmt.Fprintf(w, out)
}

func httpStatusImpl(path string, params *[]string, postData *map[string][]string) string {
	return dbStatus()
}

//
// Parse the HTTP request object into constituent parts we care about.
//
func httpParseRequest(r *http.Request) (string, *[]string, *map[string][]string) {
	path := r.URL.Path[len("/write/"):]
	query := r.URL.RawQuery
	var params []string
	if query != "" {
		params = strings.Split(query, "&")
	}
	data := (map[string][]string)(r.PostForm)
	return path, &params, &data
}

//
// Wrapper around HTTP call to write data.
//
// POST /write/foo/bar?a=b&c=d
//       DATA derf=fud&fuzz=twist
//   database = foo
//   collection = bar
//   params {a: b, c: d}
//   data = {derf: fud, fuzz: twist}
//
func httpWrite(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: /write \n", r.RemoteAddr)

	if "POST" == r.Method {
		var path string
		var params *[]string
		var data *map[string][]string
		path, params, data = httpParseRequest(r)
		out := httpWriteImpl(path, params, data)
		fmt.Fprintf(w, out)
	} else {
		// TODO: KB: [2015-09-17]: Return a proper HTTP error code.
		fmt.Fprintf(w, "Invalid HTTP method.  POST required for writes.")
	}
}

//
// Testable wrapper around dbWrite.  None of this stuff is net/http specific.  All the
// URL mapping cruft is eliminated.
//
func httpWriteImpl(path string, params *[]string, postData *map[string][]string) string {
	parts := strings.Split(path, "/")
	return dbWrite(parts[0], parts[1], params, postData)
}
