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

	//fmt.Println("YOU WANT FILTER", filter)
	// TODO: check allowedFilters and see if filter is in there
	idx := slices.Index(allowedFilters, filter)
	if idx >= 0 {
		queryValues.Add("response_status", filter)
	} else {
		return indentedJson, fmt.Errorf("filter value %s not allowed", filter)
	}

	if len(searchQuery) > 0 {
		queryValues.Add("search", searchQuery)
	}

	// TODO: add
	// response_status: all, filtered, blocked, blocked_safebrowsing, blocked_parental, whitelisted, rewritten, safe_search, processed
	// kinda works but I can pick only one and they don't quite match up with what the gui lets me pick

	//queryValues.Add("response_status", "blocked")

	// search: < arbitrary string > but it's not working, either I'm doing something wrong or it's just busted.
	//queryValues.Add("search:", "192.168.1.182")
	//queryValues.Add("search:", "newrelic.com")

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
