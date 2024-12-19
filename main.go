/*
Copyright © 2024 Eric Osborne
No header.
*/
package main

import (
	"github.com/ewosborne/adctl/cmd"
)

var version string

func main() {
	cmd.SetVersionInfo(version)
	cmd.Execute()

}
