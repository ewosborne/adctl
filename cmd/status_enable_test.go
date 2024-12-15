package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestEnable(t *testing.T) {
	// test cmd.enableCommand()

	Status, err := enableCommand()
	want := true

	if err != nil {
		t.Errorf("error testing enable: %v", err)
	}

	if Status.Protection_enabled != want {
		t.Fatal("enabling didn't return enable status")
	}

}

func TestEnableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/enable.txtar"},
	},
	)
}
