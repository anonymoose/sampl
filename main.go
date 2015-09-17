// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

var config Config

func main() {
	config := parseConfig()
	launchHttpServer(config)

	fmt.Printf(config.ListenAddr)
}
