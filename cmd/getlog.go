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

var LogArgsInstance LogArgs

func populateLogArgs(args []string) {
	if args[0] == "0" {
		args[0] = fmt.Sprintf("%v", uint32(math.MaxUint32))
	}

	LogArgsInstance = LogArgs{limit: args[0], filter: filter, search: searchQuery}
	fmt.Printf("args struct %+v\n", LogArgsInstance)
	fmt.Println("search is", LogArgsInstance.search, "done", len(LogArgsInstance.search))

}

func GetLogCmdE(cmd *cobra.Command, args []string) error {

	// TODO: change args to a struct of what it needs, passing in by args is weird.

	fmt.Println("args are", args)

	populateLogArgs(args)

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

// TODO: this is all hosed up and args should be a map or something, I think.  need to make sure I really understand how this works because it doesn't test out right.
func getLogCommand(queryLogs LogArgs) (bytes.Buffer, error) {
	var indentedJson bytes.Buffer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return indentedJson, err
	}

	baseURL.Path = "/control/querylog"

	queryValues := url.Values{}

	queryValues.Add("limit", queryLogs.limit)

	//fmt.Println("YOU WANT FILTER", filter)
	// TODO: check allowedFilters and see if filter is in there
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
