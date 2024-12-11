package cmd

import "testing"

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
