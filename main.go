// Package main provides helpers for automated testing
// with the QA Wolf Socket Go module. It contains subpackages
// like greet and others to demonstrate publishing and usage.
package main

import (
	"fmt"

	"github.com/qawolfsocket/qawolf-socket-go-module/greet"
)

func main() {
	fmt.Println(greet.Hello())
}
