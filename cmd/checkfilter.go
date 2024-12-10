/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// checkfilterCmd represents the checkfilter command
var checkfilterCmd = &cobra.Command{
	Use:   "checkfilter <string>",
	Short: "Check filters for a specific host, see if and where it's blocked. Single parameter required.",
	RunE:  CheckFilterCmdE,
}

type CheckFilterArgs struct {
	name string
}

func CheckFilterCmdE(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("need exactly one argument to checkfilter")
	}

	cfa := CheckFilterArgs{name: args[0]}

	body, err := GetFilter(cfa)
	if err != nil {
		return err
	}

	return PrintFilter(body)

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

func PrintFilter(body bytes.Buffer) error {
	fmt.Println(body.String())
	return nil
}

func init() {
	rootCmd.AddCommand(checkfilterCmd)

}
