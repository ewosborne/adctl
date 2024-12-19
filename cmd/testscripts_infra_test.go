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

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"adctl": WrapExecute,
	}))
}
