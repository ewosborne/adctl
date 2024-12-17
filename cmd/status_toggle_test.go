package cmd

import "testing"

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
