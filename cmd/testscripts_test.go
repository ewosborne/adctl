package cmd

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func setupEnv(env *testscript.Env) error {
	env.Setenv("ADCTL_HOST", os.Getenv("ADCTL_HOST"))
	env.Setenv("ADCTL_USERNAME", os.Getenv("ADCTL_USERNAME"))
	env.Setenv("ADCTL_PASSWORD", os.Getenv("ADCTL_PASSWORD"))
	return nil
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"adctl": TestscriptEntryPoint,
	}))
}
// // for new test development to run solo
// func Test_DevBox(t *testing.T) {
// 	testscript.Run(t, testscript.Params{
// 		//Dir:   "testdata/script",
// 		Setup: setupEnv,
// 		Files: []string{
// 			"testdata/script/log.txtar",

// 		},
// 	},
// 	)
// }

// all tests which can be run in parallel can go here.
func Test_AllReadonly(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{
			"testdata/script/filter_check_doubleclick.txtar",
			"testdata/script/filter_check_mit.txtar",
			"testdata/script/all_service.txtar",
			"testdata/script/rewrite_list.txtar",
			"testdata/script/log.txtar",

		},
	},
	)
}



/* everything below here are tests which read and write and so need to not be run in parallel */

func TestRewrite_AddListCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{
			"testdata/script/rewrite_add_list.txtar",
		},
	},
	)
}

func TestRewrite_DeleteCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{
			"testdata/script/rewrite_delete.txtar",
		},
	},
	)
}

func TestEnableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/service-enable.txtar"},
	},
	)
}

func TestDisableCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		//Dir:   "testdata/script",
		Setup: setupEnv,
		Files: []string{"testdata/script/service-disable.txtar"},
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
