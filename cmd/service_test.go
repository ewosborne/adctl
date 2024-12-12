package cmd

import (
	"testing"
)

func Test_ListAll(t *testing.T) {
	_, err := GetAllServices()
	if err != nil {
		t.Error(err)
	}

}
func Test_ListBlocked(t *testing.T) {
	_, err := GetBlockedServices()
	if err != nil {
		t.Error(err)
	}

}

func Test_Update(t *testing.T) {
	// TODO my APIs are weird, I'm not sure this is the right approach, skipping for now.

	t.Skip()

}
