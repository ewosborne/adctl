/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

type Status struct {
	Protection_enabled           bool
	Protection_disabled_duration uint64
}

type ReadableStatus struct {
	Protection_enabled           bool
	Protection_disabled_duration string
}

// statusCmd represents the status command
//
//lint:ignore U1000 not sure why it's unhappy
var statusGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get adblock status",
	RunE:  StatusGetCmdE,
}

func StatusGetCmdE(cmd *cobra.Command, args []string) error {
	s, err := GetStatus()
	if err != nil {
		return err
	}
	return PrintStatus(s)
}

// statusCommand prints something intelligent about what's in status
func PrintStatus(status Status) error {
	// status, err := GetStatus()
	// if err != nil {
	// 	return fmt.Errorf("error getting status: %w", err)
	// }

	var readableStatus ReadableStatus
	readableStatus.Protection_enabled = status.Protection_enabled

	if status.Protection_disabled_duration > 0 {
		readableStatus.Protection_disabled_duration = time.Duration(
			status.Protection_disabled_duration * uint64(time.Millisecond)).Truncate(time.Second).String()
		fmt.Println("rsps", readableStatus.Protection_disabled_duration)
	}

	tmp, err := json.MarshalIndent(readableStatus, "", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(tmp))
	return nil
}

func GetStatus() (Status, error) {
	var ret Status
	// get status
	// then return it?

	// build the command, it's specific to status
	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}
	baseURL.Path = "/control/status"

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	// serialize body into Status and return appropriately
	var s Status
	json.Unmarshal(body, &s)

	return s, nil
}

func init() {
	// removing this for now to see if I like it gone.
	//statusCmd.AddCommand(statusGetCmd)
}
