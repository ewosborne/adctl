/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/url"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// getlogCmd represents the getlog command
var getlogCmd = &cobra.Command{
	Use:   "getlog",
	Short: "Get logs. Optional length parameter, 0 == MaxUint32 log length.",
	RunE:  GetLogCmdE,
}

func GetLogCmdE(cmd *cobra.Command, args []string) error {

	return printLog(args)
}

func printLog(args []string) error {

	indentedJson, err := getLogCommand(args)
	if err != nil {
		return err
	}

	fmt.Println(indentedJson.String())

	return nil
}

func getLogCommand(args []string) (bytes.Buffer, error) {
	var indentedJson bytes.Buffer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return indentedJson, err
	}
	baseURL.Path = "/control/querylog"

	queryValues := url.Values{}
	if len(args) > 0 {
		if args[0] == "0" {
			queryValues.Add("limit", fmt.Sprintf("%v", uint32(math.MaxUint32)))
		} else {
			queryValues.Add("limit", args[0])
		}
	}

	baseURL.RawQuery = queryValues.Encode()

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	// if len(args) > 0 {
	// 	queryValues.Add("limit", args[0])
	// }

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return indentedJson, err
	}

	json.Indent(&indentedJson, body, "", "  ")

	return indentedJson, nil
}

func init() {
	rootCmd.AddCommand(getlogCmd)
}
