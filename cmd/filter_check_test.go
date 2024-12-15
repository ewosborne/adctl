package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func Test_Checkfilter(t *testing.T) {
	cfa := CheckFilterArgs{name: "www.doubleclick.net"}
	_, err := GetFilter(cfa)
	if err != nil {
		t.Errorf("error in GetFilter: %v", err)
	}
}

// I could run any parallel read-only tests in here I suppose
func TestFilterCheck(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{
			"testdata/script/filter_check_doubleclick.txtar",
			"testdata/script/filter_check_mit.txtar",
		},
	},
	)
}
