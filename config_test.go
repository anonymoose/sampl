package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseConfig(t *testing.T) {
	conf := parseConfig()
	assert.Equal(t, DEFAULT_LISTEN_ADDR, conf.ListenAddr)
	assert.Equal(t, DEFAULT_LISTEN_PORT, conf.ListenPort)
}
