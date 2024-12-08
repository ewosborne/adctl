package cmd

import (
	"encoding/json"
	"testing"
)

func Test_Getlog(t *testing.T) {
	// test only a small log thing
	log, err := getLogCommand([]string{"10"})

	if err != nil {
		t.Error("error getting getLogCommand", err)
	}

	// would be nice to test for number of entries but I don't know how.
	if !json.Valid(log.Bytes()) {
		t.Error("invalid json log", err)
	}
}
