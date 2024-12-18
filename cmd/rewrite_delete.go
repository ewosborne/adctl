/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// rewriteDeleteCmd represents the delete command
var rewriteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a rewrite",
	RunE:  RewriteDeleteCmdE,
}

func RewriteDeleteCmdE(cmd *cobra.Command, args []string) error {
	//  takes a domain and an answer. API requires both. I could get clever about matching on existing domain and deleting its answer but $$LATER. TODO.

	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	answer, err := cmd.Flags().GetString("answer")
	if err != nil {
		return err
	}

	err = deleteRewrite(domain, answer)
	if err != nil {
		return err
	}
	printRewriteList()
	return nil
}

func deleteRewrite(domain string, answer string) error {

	var requestBody = make(map[string]any)
	requestBody["domain"] = domain
	requestBody["answer"] = answer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return err
	}

	baseURL.Path = "/control/rewrite/delete"

	enableQuery := common.CommandArgs{
		Method:      "POST",
		URL:         baseURL,
		RequestBody: requestBody,
	}

	_, err = common.SendCommand(enableQuery)
	if err != nil {
		return err
	}

	return nil
}

// TODO add '--all' maybe?  to delete everything?
func init() {
	rewriteCmd.AddCommand(rewriteDeleteCmd)
	rewriteDeleteCmd.Flags().String("domain", "", "Name or wildcard to match on")
	rewriteDeleteCmd.Flags().String("answer", "", "Answer to supply in response. IP address, domain name, or some weird special stuff around A and AAAA.")
}
