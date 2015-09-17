// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseConfig(t *testing.T) {
	conf := parseConfig()
	assert.Equal(t, DEFAULT_LISTEN_ADDR, conf.ListenAddr)
	assert.Equal(t, DEFAULT_LISTEN_PORT, conf.ListenPort)
	assert.Equal(t, DEFAULT_CLIENT_ADDR, conf.ClientAddr)
	assert.Equal(t, DEFAULT_CLIENT_PORT, conf.ClientPort)
}
