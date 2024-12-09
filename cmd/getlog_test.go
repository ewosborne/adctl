package cmd

import (
	"encoding/json"
	"fmt"
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

func Test_Getlog_Filter(t *testing.T) {
	// test with an allowed and disallowed filter.
	initialFilter := filter
	filter = "all"
	var err error
	_, err = getLogCommand([]string{"10"})

	if err != nil {
		t.Error("error getting getLogCommand with valid filter", err)
	}

	filter = "bogon"
	_, err = getLogCommand([]string{"10"})

	if err == nil {
		t.Error("tried getLogCommand with invalid filter and didn't get error")
	}

	filter = initialFilter // reset because this would otherwise carry across tests

}
func Test_Getlog_Search(t *testing.T) {
	// TODO shoudl really be more comprehensive than just this

	searchQuery = "netflix.com"
	var err error
	_, err = getLogCommand([]string{"10"})

	if err != nil {
		t.Error("got non-fill error testing getLogCommand")
		fmt.Println(err)
	}
}
