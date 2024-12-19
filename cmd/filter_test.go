package cmd

import (
	"testing"
)

func TestFilterCheck(t *testing.T) {
	cfa := CheckFilterArgs{name: "www.doubleclick.net"}
	_, err := GetFilter(cfa)
	if err != nil {
		t.Errorf("error in GetFilter: %v", err)
	}
}
