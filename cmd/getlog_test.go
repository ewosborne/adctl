package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type ValidQueryResult struct {
	Oldest string `json:"oldest"`
}

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
	body, err := getLogCommand([]string{"10"})

	if err != nil {
		t.Error("error getting getLogCommand with valid filter", err)
	}

	present, err := checkBufferForJson(body)
	if err != nil {
		t.Error("got non-fill error testing getLogCommand")
		fmt.Println(err)
	}
	if !present {
		t.Error("'oldest' not present in json")
	}

	filter = "bogon"
	_, err = getLogCommand([]string{"10"}) // do not capture body here, it's empty, I just care about error
	filter = initialFilter                 // reset because this would otherwise carry across tests

	if err == nil {
		t.Error("tried getLogCommand with invalid filter and didn't get error")
	}

}
func Test_Getlog_Search(t *testing.T) {

	var searchQuery string = "example.com" // doesn't matter what goes here, even with empty results I still get valid json
	var err error
	body, err := getLogCommand([]string{"10"})

	if err != nil {
		t.Error("got non-fill error testing getLogCommand")
		fmt.Println(err)
	}

	present, err := checkBufferForJson(body)
	if err != nil {
		t.Error("got non-fill error testing getLogCommand")
		fmt.Println(err)
	}
	if !present {
		t.Error("'oldest' not present in json")
	}
}

func checkBufferForJson(body bytes.Buffer) (bool, error) {
	// passed in []bytes
	// marshal it and return whether 'oldest' is presetn
	// fail otherwise

	//var err error
	//tmpJson := make(map[string]string)
	tmpJson := ValidQueryResult{}

	err := json.Unmarshal(body.Bytes(), &tmpJson)
	if err != nil {
		return false, fmt.Errorf("can't unmarshal json")
	}

	if len(tmpJson.Oldest) < 1 {
		return false, fmt.Errorf("json.oldest is weird")
	}
	return true, nil

}
