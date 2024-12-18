package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

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
