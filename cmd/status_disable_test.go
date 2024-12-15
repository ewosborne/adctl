package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestDisableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/disable.txtar"},
	},
	)
}

func TestDisableTimeoutCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/disable-timeout.txtar"},
	},
	)
}

func TestDisable_Permanent(t *testing.T) {
	// test cmd.enableCommand()

	dTime := DisableTime{HasTimeout: false}

	Status, err := disableCommand(dTime)
	want := false

	if err != nil {
		t.Errorf("error testing enable: %v", err)
	}

	if Status.Protection_enabled != want {
		t.Fatalf("disabling didn't return enable status")
	}

}

func TestDisable_Temporary(t *testing.T) {
	// test cmd.enableCommand()

	dTime := DisableTime{HasTimeout: true, Duration: "30s"}

	Status, err := disableCommand(dTime)
	want := false

	if err != nil {
		t.Errorf("error testing enable: %v", err)
	}

	if Status.Protection_enabled != want {
		t.Fatalf("disabling didn't return enable status")
	}

	// protection enable time is in msec, test is 30sec, 25sec should be fine
	if Status.Protection_disabled_duration < 25000 || Status.Protection_disabled_duration > 30000 {
		t.Fatal("protection status isn't in the 25-30sec range", Status.Protection_disabled_duration)
	}
}
