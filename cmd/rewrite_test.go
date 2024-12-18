package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestRewriteListCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{"testdata/script/rewrite_list.txtar"},
	},
	)
}

func TestRewriteDeleteCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{"testdata/script/rewrite_delete.txtar"},
	},
	)
}

func TestRewriteAddCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{"testdata/script/rewrite_add.txtar"},
	},
	)
}
