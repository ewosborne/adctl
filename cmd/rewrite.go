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

// rewriteCmd represents the rewrite command
var rewriteCmd = &cobra.Command{
	Use:   "rewrite",
	Short: "Control DNS rewrites",
	Long:  "Add, delete, or list DNS rewrites.",
}

var rewriteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS rewrites",
	RunE:  RewriteListCmdE,
}

var rewriteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a rewrite",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RewriteCommand(cmd, args, true)
	},
}

var rewriteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a rewrite",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RewriteCommand(cmd, args, false)
	},
}

func init() {
	rootCmd.AddCommand(rewriteCmd)
	rewriteCmd.AddCommand(rewriteListCmd)

	rewriteCmd.AddCommand(rewriteAddCmd)
	rewriteCmd.AddCommand(rewriteDeleteCmd)

	rewriteAddCmd.Flags().String("domain", "", "Name or wildcard to match on")
	rewriteAddCmd.Flags().String("answer", "", "Answer to supply in response. IP address, domain name, or some weird special stuff around A and AAAA.")

	rewriteDeleteCmd.Flags().String("domain", "", "Name or wildcard to match on")
	rewriteDeleteCmd.Flags().String("answer", "", "Answer to supply in response. IP address, domain name, or some weird special stuff around A and AAAA.")

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

func RewriteCommand(cmd *cobra.Command, args []string, add bool) error {

	// if add is true then add
	// if add is false then delete

	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	answer, err := cmd.Flags().GetString("answer")
	if err != nil {
		return err
	}

	err = doRewriteAction(domain, answer, add)
	if err != nil {
		return err
	}
	printRewriteList()
	return nil

}

func doRewriteAction(domain string, answer string, add bool) error {

	var requestBody = make(map[string]any)
	var err error
	requestBody["domain"] = domain
	requestBody["answer"] = answer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return err
	}

	switch add {
	case true:
		baseURL.Path = "/control/rewrite/add"
	case false:
		baseURL.Path = "/control/rewrite/delete"
	}

	enableQuery := common.CommandArgs{
		Method:      "POST",
		URL:         baseURL,
		RequestBody: requestBody,
	}

	if add {
		// delete before adding because adding isn't idempotent.
		err = doRewriteAction(domain, answer, false)
		if err != nil {
			return err
		}
	}

	_, err = common.SendCommand(enableQuery)
	if err != nil {
		return err
	}

	return nil
}
