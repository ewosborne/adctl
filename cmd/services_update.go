/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// serviceUpdateCmd represents the update command
var serviceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Block or unblock one or more services",
	RunE:  UpdateServiceCmdE,
}

// populated as flags, see init()
// TODO: put these in a struct, clean them up?
var toPermit []string
var toBlock []string
var permitAll bool

type ServiceLists struct {
	permit []string
	block  []string
}

func unique(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func UpdateServiceCmdE(cmd *cobra.Command, args []string) error {

	// TODO hack
	if len(toBlock) == 0 && len(toPermit) == 0 {
		return fmt.Errorf("need either permit or blocked flag")
	}
	// first tidy up.  should I wrap these in a struct? TODO
	toBlock = unique(toBlock)
	toPermit = unique(toPermit)

	svcs := ServiceLists{block: toBlock, permit: toPermit}

	// should I pass in toPermit and toBlock or leave them global here?
	//   does passing them as args make testing easier? TODO
	err := updateServices(svcs)
	if err != nil {
		return fmt.Errorf("error updating services %w", err)
	}

	return nil
}

func updateServices(svcs ServiceLists) error {

	// note that blocked has a schedule as well.  That just gets passed transparently through, I don't touch it.
	blocked, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error calling GetBlockedServices %w", err)
	}

	// take the list of currently blocked services, may be none.
	//  add all new toBlock and then remove all toPermit
	tmp := make(map[string]bool)
	for _, s := range blocked.IDs {
		tmp[s] = true
	}
	for _, s := range svcs.block {
		tmp[s] = true
	}

	for _, s := range svcs.permit {
		tmp[s] = false
	}

	// turn service map back into the final list of services
	newlist := []string{}
	for k, v := range tmp {
		if v {
			newlist = append(newlist, k)
		}
	}

	// TODO: I left this line out once and tests didn't catch it.  need a test.
	blocked.IDs = newlist
	// // special case
	// if permitAll {
	// 	newlist = []string{}
	// }

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
	// TODO: check to make sure that what we just pushed looks like what the server thinks
	s, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error getting blocked services %w", err)
	}

	slices.Sort(blocked.IDs)
	slices.Sort(s.IDs)
	s.IDs = []string{"foobar"}
	if !slices.Equal(blocked.IDs, s.IDs) {
		return fmt.Errorf("service lists unequal: %v %v", blocked.IDs, s.IDs)
	}

	return nil

}

func init() {
	servicesCmd.AddCommand(serviceUpdateCmd)
	serviceUpdateCmd.Flags().StringSliceVar(&toPermit, "permit", []string{}, "CSV of services to permit")
	serviceUpdateCmd.Flags().StringSliceVar(&toBlock, "block", []string{}, "CSV of services to block")
	serviceUpdateCmd.Flags().BoolVarP(&permitAll, "permit-all", "", false, "Permit all services")

}
