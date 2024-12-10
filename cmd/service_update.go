/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// updateServiceCmd represents the block command
var updateServiceCmd = &cobra.Command{
	Use:   "update",
	Short: "Block or unblock one or more services",
	RunE:  UpdateServiceCmdE,
}

var toPermit []string
var toBlock []string

func unique(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func UpdateServiceCmdE(cmd *cobra.Command, args []string) error {

	// first tidy up
	toBlock = unique(toBlock)
	toPermit = unique(toPermit)

	// fmt.Println("to block", toBlock)
	// fmt.Println("to permit", toPermit)

	err := updateServices()
	if err != nil {
		return fmt.Errorf("error updating services %w", err)
	}

	return nil
}

// TODO: rethink this.  maybe prepend '+' to services to block and '-' to unblock, do it in one command? also maybe support 'all'?
//
//	also TODO: think about service filter.  maybe just take what's there and keep it but don't change via cli. yeah, let's do that.
//     or --block and --unblock but allow multiples? that seems more idiomatic.

func updateServices() error {

	blocked, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error calling GetBlockedServices %w", err)
	}

	// fmt.Println("currently blocked", blocked.IDs)
	// fmt.Println("you want to block", toBlock)
	// fmt.Println("you want to permit", toPermit)

	// ...because I have no set primitives...
	tmp := make(map[string]bool)
	for _, s := range blocked.IDs {
		tmp[s] = true
	}
	for _, s := range toBlock {
		tmp[s] = true
	}
	for _, s := range toPermit {
		tmp[s] = false
	}

	newlist := []string{}
	for k, v := range tmp {
		if v == true {
			newlist = append(newlist, k)
		}
	}

	blocked.IDs = unique(newlist)

	// TODO: what about schedule?  skip for now but what if it comes down with a schedule and I don't push one up?

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return err
	}

	baseURL.Path = "/control/blocked_services/update"

	var requestBody = make(map[string]any)
	requestBody["ids"] = blocked.IDs
	requestBody["schedule"] = blocked.Schedule

	//fmt.Println("going to update with", requestBody)

	// put it all together
	enableQuery := common.CommandArgs{
		Method:      "PUT",
		URL:         baseURL,
		RequestBody: requestBody,
	}

	// don't care about body here
	_, err = common.SendCommand(enableQuery)
	if err != nil {
		return err
	}

	return nil

}

func init() {
	serviceCmd.AddCommand(updateServiceCmd)
	updateServiceCmd.Flags().StringSliceVar(&toPermit, "permit", []string{}, "CSV of services to permit")
	updateServiceCmd.Flags().StringSliceVar(&toBlock, "block", []string{}, "CSV of services to block")

}
