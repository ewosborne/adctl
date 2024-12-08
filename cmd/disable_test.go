package cmd

import "testing"

func TestDisable_Permanent(t *testing.T) {
	// test cmd.enableCommand()

	Status, err := disableCommand([]string{})
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

	Status, err := disableCommand([]string{"30s"})
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
