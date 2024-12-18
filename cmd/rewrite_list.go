/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS rewrites",
	RunE:  RewriteListCmdE,
}

type RewriteList []map[string]string

func RewriteListCmdE(cmd *cobra.Command, args []string) error {
	return printRewriteList()
}

func printRewriteList() error {

	status, err := rewriteListCommand()
	if err != nil {
		return err
	}

	tmp, err := json.MarshalIndent(status, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(tmp))
	return nil
}

func rewriteListCommand() (RewriteList, error) {

	//var ret = make(map[string]any)
	var ret RewriteList

	// list is a GET, takes no params
	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}
	baseURL.Path = "/control/rewrite/list"

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func init() {
	rewriteCmd.AddCommand(listCmd)
}
