package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type CommandArgs struct {
	RequestBody map[string]any
	Method      string
	URL         url.URL
}

func GetBaseURL() (url.URL, error) {
	var ret = url.URL{Scheme: "http"}

	h, err := getHost()
	if err != nil {
		return ret, err
	}

	ret.Host = fmt.Sprintf(h)
	return ret, nil

}

func getHost() (string, error) {
	ret, present := os.LookupEnv("ADCTL_HOST")
	if !present {
		return "", fmt.Errorf("can't find ADCTL_HOST")
	}
	return ret, nil
}

func AbleCommand(state bool, durationString string) error {
	//log.Print("in AbleCommand with duration ", durationString)

	// base url
	baseURL, err := GetBaseURL()
	if err != nil {
		return err
	}

	baseURL.Path = "/control/protection"

	// data for post
	var requestBody = make(map[string]any)
	requestBody["enabled"] = state

	var duration uint64
	if len(durationString) > 0 { // is this ugly?

		tmp, err := time.ParseDuration(durationString)
		if err != nil {
			return fmt.Errorf("time.ParseDuration: %w", err)
		}
		duration = uint64(tmp.Milliseconds())
	}

	requestBody["duration"] = duration

	// put it all together
	enableQuery := CommandArgs{
		Method:      "POST",
		URL:         baseURL,
		RequestBody: requestBody,
	}

	// don't care about body here
	_, err = SendCommand(enableQuery)
	if err != nil {
		return err
	}

	return nil
}

func SendCommand(ca CommandArgs) ([]byte, error) {
	//log.Print("in SendCommand")
	//log.Print("need to call ", ca.URL.String())

	//var client *http.Client

	var jsonData []byte
	var err error

	// turn params into json.  not sure if I can safely do this to all verbs.
	if ca.Method == "POST" {
		jsonData, err = json.Marshal(ca.RequestBody)
		if err != nil {
			return nil, fmt.Errorf("error marshaling json: %w", err)
		}
	}

	// create the final request
	request, err := http.NewRequest(ca.Method, ca.URL.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// set request headers
	request.Header.Set("Content-Type", "application/json")

	// set basic auth TODO put this somewhere smart.
	un, present := os.LookupEnv("ADCTL_USERNAME")
	if !present {
		return nil, fmt.Errorf("can't find ADCTL_USERNAME")
	}
	pw, present := os.LookupEnv("ADCTL_PASSWORD")
	if !present {
		return nil, fmt.Errorf("can't find ADCTL_PASSWORD")
	}
	request.SetBasicAuth(un, pw)

	// connect.  TODO old stuff let me set timeouts to handle short dns timeouts and
	//   long log fetches.  bother with it here?
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error Do'ing request: %w", err)
	}
	defer resp.Body.Close()

	// read response
	// Read response but really I just want to know if there's an error
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v\n", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status code not 200: %v\n", resp.Status)
	}

	return body, nil
}
