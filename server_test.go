// Copyright 2015 Palm Valley Data Lab. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serverStatus(t *testing.T) {
	statusMsg := httpStatusImpl("", nil, nil)
	assert.Equal(t, "Status: OK", statusMsg)

}

func Test_writeImpl(t *testing.T) {
	writeMsg := httpWriteImpl("foo/bar", nil, nil)
	assert.Equal(t, "Write: OK", writeMsg)
}
