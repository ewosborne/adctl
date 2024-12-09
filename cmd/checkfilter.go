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
	fmt.Printf("checkfilter called %+v\n", cfa)

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
		return indentedJson, err
	}
	baseURL.Path = "/control/filtering/check_host"
	queryValues := url.Values{}
	queryValues.Add("name", cfa.name)

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

	// serialize body into Status and return appropriately
	json.Unmarshal(body, &s)

	return s, nil

	return indentedJson, nil
}

func PrintFilter(bytes.Buffer) error {
	// unpack into json and print
	return nil
}

func init() {
	rootCmd.AddCommand(checkfilterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkfilterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkfilterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
