package cmd

import (
	"testing"
)

// TODO: testscript which just looks for valid json
func TestToggle(t *testing.T) {
	// get initial state
	// toggle to other, make sure it sticks
	// toggle back

	initialState, err := GetStatus()
	if err != nil {
		t.Errorf("error getting initial status: %v", err)
	}

	err = toggleCommand()
	if err != nil {
		t.Errorf("error toggling command: %v", err)
	}

	secondState, err := GetStatus()
	if err != nil {
		t.Errorf("error getting second status: %v", err)
	}

	// at this point they should be opposite so equality is an error
	if initialState.Protection_enabled == secondState.Protection_enabled {
		t.Errorf("first toggle: protection states do not match")
	}

	err = toggleCommand()
	if err != nil {
		t.Errorf("error toggling command: %v", err)
	}
	thirdState, err := GetStatus()
	if err != nil {
		t.Errorf("error getting third status: %v", err)
	}

	// at this point they should be equal again so inequality is an error
	if initialState.Protection_enabled != thirdState.Protection_enabled {
		t.Errorf("second toggle: protection states match but shouldn't: %v %v", initialState.Protection_enabled, thirdState.Protection_enabled)
	}

}

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
