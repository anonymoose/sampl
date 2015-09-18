// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func chkerr(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
		panic(e)
	}
}
