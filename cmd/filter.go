/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ewosborne/adctl/common"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Check filter for entities",
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check filters for a specific host, see if and where it's blocked. Single parameter required.",
	Long:  `long help TBD`,
	RunE:  CheckFilterCmdE,
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.AddCommand(checkCmd)

}

type CheckFilterArgs struct {
	name string
}

func CheckFilterCmdE(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("need exactly one argument to checkfilter")
	}

	cfa := CheckFilterArgs{name: args[0]}

	return PrintFilter(cfa)

}

func PrintFilter(cfa CheckFilterArgs) error {
	body, err := GetFilter(cfa)
	if err != nil {
		return err
	}
	fmt.Println(body.String())
	return nil
}

func GetFilter(cfa CheckFilterArgs) (bytes.Buffer, error) {

	var ret bytes.Buffer
	// get the filter
	// pack it as json

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}
	baseURL.Path = "/control/filtering/check_host"
	queryValues := url.Values{}
	queryValues.Add("name", cfa.name)

	baseURL.RawQuery = queryValues.Encode()

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	json.Indent(&ret, body, "", "  ")

	return ret, nil
}
