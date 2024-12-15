package cmd

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func setupEnv(env *testscript.Env) error {
	env.Setenv("ADCTL_HOST", os.Getenv("ADCTL_HOST"))
	env.Setenv("ADCTL_USERNAME", os.Getenv("ADCTL_USERNAME"))
	env.Setenv("ADCTL_PASSWORD", os.Getenv("ADCTL_PASSWORD"))
	return nil
}

// note: need to run all these seperately because I think there's some
//
//		concurrency if I run them all via Run:Dir: and I don't know how to
//	 turn it off
func TestDisableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/disable.txtar"},
	},
	)
}

func TestEnableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/enable.txtar"},
	},
	)
}
