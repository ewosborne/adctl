package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CommandArgs struct {
	RequestBody map[string]any
	Method      string
	URL         url.URL
}

// LoadConfig reads configuration files and sets environment variables if they're not already set
func LoadConfig() error {
	configs := map[string]string{}

	if err := readConfigFile("/etc/adctl.conf", configs); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error reading /etc/adctl.conf: %w", err)
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home directory: %w", err)
	}
	userConfigPath := filepath.Join(homeDir, ".adctl")
	if err := readConfigFile(userConfigPath, configs); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error reading %s: %w", userConfigPath, err)
		}
	}

	envVarsSet := 0
	for key, value := range configs {
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
			envVarsSet++
		}
	}

	return nil
}

// readConfigFile reads a configuration file and parses key=value pairs
func readConfigFile(filename string, configs map[string]string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		if key == "ADCTL_HOST" || key == "ADCTL_USERNAME" || key == "ADCTL_PASSWORD" {
			configs[key] = value
		}
	}

	return scanner.Err()
}

func GetBaseURL() (url.URL, error) {
	h, err := getHost()
	if err != nil {
		return url.URL{}, err
	}

	scheme := "http"
	if len(h) >= 8 && h[:8] == "https://" {
		scheme = "https"
		h = h[8:]
	} else if len(h) >= 7 && h[:7] == "http://" {
		scheme = "http"
		h = h[7:]
	}
	ret := url.URL{Scheme: scheme, Host: fmt.Sprint(h)}
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
	if ca.Method == "POST" || ca.Method == "PUT" {
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

	// set basic auth
	un, present := os.LookupEnv("ADCTL_USERNAME")
	if !present {
		return nil, fmt.Errorf("can't find ADCTL_USERNAME")
	}
	pw, present := os.LookupEnv("ADCTL_PASSWORD")
	if !present {
		return nil, fmt.Errorf("can't find ADCTL_PASSWORD")
	}
	request.SetBasicAuth(un, pw)

	// TODO: debug flag for stuff like this.
	// if request.Method == "GET" {
	// 	fmt.Printf("sending request %+v\n", request)
	// }

	// connect.  Old implementation let me set timeouts to handle short dns timeouts and
	//   long log fetches.  bother with it here? skipping for now.
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
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code not 200: %v", resp.Status)
	}

	return body, nil
}
