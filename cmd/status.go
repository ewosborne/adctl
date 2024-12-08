/*
Copyright © 2024 Eric Osborne
No header.
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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get adblock status",
	RunE:  StatusCmdE,
}

func StatusCmdE(cmd *cobra.Command, args []string) error {
	s, err := GetStatus()
	if err != nil {
		return err
	}
	return PrintStatus(s)
}

// statusCommand calls getStatus and prints something
func PrintStatus(status Status) error {
	// status, err := GetStatus()
	// if err != nil {
	// 	return fmt.Errorf("error getting status: %w", err)
	// }

	var statusString string
	if status.Protection_enabled {
		statusString = "enabled"
	} else {
		statusString = "disabled"
	}

	if status.Protection_disabled_duration > 0 {
		remaining := time.Duration(
			status.Protection_disabled_duration * uint64(time.Millisecond)).Truncate(time.Second)

		statusString += fmt.Sprintf(" for %v", remaining)
	}

	fmt.Printf("adguard is %v\n", statusString)

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
	rootCmd.AddCommand(statusCmd)

}
