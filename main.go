package main

import (
	"fmt"
)

var config Config

func main() {
	config := parseConfig()

	fmt.Printf(config.ListenAddr)
}
