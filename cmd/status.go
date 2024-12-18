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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check and change adblocking status",
	RunE:  StatusGetCmdE,
}

var statusToggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle adblocker between enabled and disabled.",
	RunE:  ToggleCmdE,
}

var statusEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable ad blocking",
	RunE:  StatusEnableCmdE,
}

// statusCmd represents the status command
//
//lint:ignore U1000 not sure why it's unhappy
var statusGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get adblock status",
	RunE:  StatusGetCmdE,
}

var statusDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable ad blocker. Optional duration in time.Duration format.",
	Args:  cobra.RangeArgs(0, 1),
	RunE:  StatusDisableCmdE,
}

type Status struct {
	Protection_enabled           bool
	Protection_disabled_duration uint64
}

type DisableTime struct {
	Duration   string
	HasTimeout bool
}

type ReadableStatus struct {
	Protection_enabled           bool
	Protection_disabled_duration string
}

func ToggleCmdE(cmd *cobra.Command, args []string) error {
	return printToggle()
}

func StatusEnableCmdE(cmd *cobra.Command, flags []string) error {
	return printEnable()
}

func StatusGetCmdE(cmd *cobra.Command, args []string) error {
	s, err := GetStatus()
	if err != nil {
		return err
	}
	return PrintStatus(s)
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// left this one out on purpose, not sure why TODO figure it out
	//statusCmd.AddCommand(statusGetCmd)
	statusCmd.AddCommand(statusToggleCmd)
	statusCmd.AddCommand(statusEnableCmd)

	statusCmd.AddCommand(statusDisableCmd)

}

func printToggle() error {
	var err error

	err = toggleCommand()
	if err != nil {
		return err
	}

	status, err := GetStatus()
	if err != nil {
		return err
	}
	PrintStatus(status)
	return nil
}

func toggleCommand() error {
	status, err := GetStatus()
	if err != nil {
		return err
	}

	dTime := DisableTime{HasTimeout: false}
	switch status.Protection_enabled {
	case true:
		disableCommand(dTime)
	case false:
		enableCommand()
	}

	return nil
}

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

func printEnable() error {

	var err error
	status, err := enableCommand()
	if err != nil {
		return err
	}
	err = PrintStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func enableCommand() (Status, error) {

	err := common.AbleCommand(true, "")
	if err != nil {
		return Status{}, err
	}

	status, err := GetStatus()
	if err != nil {
		return Status{}, err
	}

	return status, nil
}

func StatusDisableCmdE(cmd *cobra.Command, args []string) error {

	var dTime = DisableTime{}

	switch len(args) {
	case 0:
		dTime.HasTimeout = false
	case 1:
		dTime.HasTimeout = true
		dTime.Duration = args[0]
	default:
		return fmt.Errorf("only one arg allowed for disable")
	}

	return printDisable(dTime)
}

func printDisable(dTime DisableTime) error {

	var err error

	status, err := disableCommand(dTime)
	if err != nil {
		return err
	}

	PrintStatus(status)

	return err

}

func disableCommand(dTime DisableTime) (Status, error) {

	var err error

	if dTime.HasTimeout {
		err = common.AbleCommand(false, dTime.Duration)
	} else {
		err = common.AbleCommand(false, "")
	}

	if err != nil {
		return Status{}, err
	}

	s, err := GetStatus()

	return s, err
}
