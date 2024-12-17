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

var TestLogArgsInstance = LogArgs{limit: "10", filter: "all"}

// TODO: tescript test.  maybe 'adctl log get 5' then validate length and json and that there's an 'oldest'?

func Test_LogGet(t *testing.T) {
	// test only a small log thing
	log, err := getLogCommand(TestLogArgsInstance)

	if err != nil {
		t.Error("error getting getLogCommand", err)
	}

	if !json.Valid(log.Bytes()) {
		t.Error("invalid json log", err)
	}
}

func Test_LogGet_Filter(t *testing.T) {

	// test with an allowed and disallowed filter.
	// filter comes from the variable declared as a flag
	var err error
	body, err := getLogCommand(TestLogArgsInstance)

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

	initialFilter := TestLogArgsInstance.filter
	TestLogArgsInstance.filter = "bogon"
	_, err = getLogCommand(TestLogArgsInstance) // do not capture body here, it's empty, I just care about error
	TestLogArgsInstance.filter = initialFilter  // reset because this would otherwise carry across tests

	if err == nil {
		t.Error("tried getLogCommand with invalid filter and didn't get error")
	}

}
func Test_LogGet_Search(t *testing.T) {

	var err error

	TestLogArgsInstance.search = "example.com"
	body, err := getLogCommand(TestLogArgsInstance)

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

	tmpJson := ValidQueryResult{}

	err := json.Unmarshal(body.Bytes(), &tmpJson)
	if err != nil {
		return false, fmt.Errorf("can't unmarshal json")
	}

	if len(tmpJson.Oldest) < 1 {
		return false, fmt.Errorf("json.oldest is empty")
	}
	return true, nil

}
