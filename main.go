/*
Copyright © 2024 Eric Osborne
No header.
*/
package main

import (
	"os"

	"github.com/ewosborne/adctl/cmd"
)

var version string

func main() {
	os.Exit(Main())

}

// doing this to make testscript happy
func Main() int {
	cmd.SetVersionInfo(version)
	cmd.Execute()
	return 0
}
