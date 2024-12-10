package cmd

import "testing"

func Test_Checkfilter(t *testing.T) {
	cfa := CheckFilterArgs{name: "www.doubleclick.net"}
	_, err := GetFilter(cfa)
	if err != nil {
		t.Errorf("error in GetFilter: %v", err)
	}
}
