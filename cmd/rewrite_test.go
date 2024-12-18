package cmd

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestRewrite_AddListDeleteCLI(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Setup: setupEnv,
		Files: []string{"testdata/script/rewrite_list.txtar"},
	},
	)
}
