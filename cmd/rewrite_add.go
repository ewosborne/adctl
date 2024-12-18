/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// rewriteAddCmd represents the add command
var rewriteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a rewrite",
	RunE:  RewriteAddCmdE,
}

func RewriteAddCmdE(cmd *cobra.Command, args []string) error {
	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	answer, err := cmd.Flags().GetString("answer")
	if err != nil {
		return err
	}

	err = addRewrite(domain, answer)
	if err != nil {
		return err
	}
	printRewriteList()
	return nil
}

// TODO this is 99% the same as deleteRewrite, combine them.
func addRewrite(domain string, answer string) error {

	var requestBody = make(map[string]any)
	var err error
	requestBody["domain"] = domain
	requestBody["answer"] = answer

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return err
	}

	baseURL.Path = "/control/rewrite/add"

	enableQuery := common.CommandArgs{
		Method:      "POST",
		URL:         baseURL,
		RequestBody: requestBody,
	}

	// delete before adding because adding isn't idempotent.  TODO handle this better?
	err = deleteRewrite(domain, answer)
	if err != nil {
		return err
	}

	_, err = common.SendCommand(enableQuery)
	if err != nil {
		return err
	}

	// TODO: print result.
	//printRewriteList()

	return nil

}

func init() {
	rewriteCmd.AddCommand(rewriteAddCmd)
	rewriteAddCmd.Flags().String("domain", "", "Name or wildcard to match on")
	rewriteAddCmd.Flags().String("answer", "", "Answer to supply in response. IP address, domain name, or some weird special stuff around A and AAAA.")
}
