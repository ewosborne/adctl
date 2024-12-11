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
	"slices"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// getlogCmd represents the getlog command
var getlogCmd = &cobra.Command{
	Use:   "getlog",
	Short: "Get logs. Optional length parameter, 0 == MaxUint32 log length.",
	RunE:  GetLogCmdE,
}

var filter string
var searchQuery string
var allowedFilters = []string{
	"all", "filtered", "blocked",
	"blocked_safebrowsing", "blocked_parental",
	"whitelisted", "rewritten", "safe_search", "processed",
}

// LogArgs is what's passed around with arguments to the log query
type LogArgs struct {
	limit  string
	filter string
	search string
}

func GetLogCmdE(cmd *cobra.Command, args []string) error {

	//populateLogArgs(args)
	LogArgsInstance := LogArgs{filter: filter, search: searchQuery}

	// if there are no args then do nothing
	// if there's one arg it's either zero or it's not
	//   if it's zero then bump it up but in a platform-specific way

	switch len(args) {
	case 0:
		// do nothing
	case 1:
		if args[0] == "0" {
			args[0] = fmt.Sprintf("%v", uint32(math.MaxUint32))
		}
		LogArgsInstance.limit = args[0]
	default:
		return fmt.Errorf("too many args to GetLogCmdE: %v", len(args))
	}

	return printLog(LogArgsInstance)
}

func printLog(queryLogs LogArgs) error {

	indentedJson, err := getLogCommand(queryLogs)
	if err != nil {
		return err
	}

	fmt.Println(indentedJson.String())

	return nil
}

func getLogCommand(queryLogs LogArgs) (bytes.Buffer, error) {
	var indentedJson bytes.Buffer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return indentedJson, err
	}

	baseURL.Path = "/control/querylog"

	queryValues := url.Values{}

	queryValues.Add("limit", queryLogs.limit)

	idx := slices.Index(allowedFilters, queryLogs.filter)
	if idx >= 0 {
		queryValues.Add("response_status", queryLogs.filter)
	} else {
		return indentedJson, fmt.Errorf("filter value %s not allowed", queryLogs.filter)
	}

	if len(queryLogs.search) > 0 {
		queryValues.Add("search", queryLogs.search)
	}

	baseURL.RawQuery = queryValues.Encode()

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return indentedJson, err
	}

	json.Indent(&indentedJson, body, "", "  ")

	return indentedJson, nil
}

func init() {
	rootCmd.AddCommand(getlogCmd)
	getlogCmd.Flags().StringVarP(&filter, "filter", "", "all", fmt.Sprintf("one of: %#v", allowedFilters))
	getlogCmd.Flags().StringVarP(&searchQuery, "search", "", "", "string to search for in logs.")

}
