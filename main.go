// Package main provides helpers for automated testing
// with the QA Wolf Socket Go module. It contains subpackages
// like qawolf_package and others to demonstrate publishing and usage.
package main

import (
	"fmt"

	"github.com/qawolfsocket/qawolf-socket-go-module/qawolf_package"
)

func main() {
	fmt.Println(qawolf_package.Hello())
}
