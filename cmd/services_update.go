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
var toPermit []string
var toBlock []string

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

	// should I pass in toPermit and toBlock or leave them global here?
	//   does passing them as args make testing easier? TODO
	err := updateServices()
	if err != nil {
		return fmt.Errorf("error updating services %w", err)
	}

	return nil
}

func updateServices() error {

	// note that blocked has a schedule as well.  That just gets passed transparently through, I don't touch it.
	blocked, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error calling GetBlockedServices %w", err)
	}

	// take the list of currently blocked services, may be none.
	//  add all new toBlock and then remove all toPermit
	tmp := make(map[string]bool)

	// start with what we're already blocking
	for _, s := range blocked.IDs {
		tmp[s] = true
	}

	// add new stuff to block
	for _, s := range svcs.block {
		tmp[s] = true
	}

	// subtract anything to unblock
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
	blocked.IDs = newlist

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
<<<<<<< Updated upstream
=======
	// TODO: check to make sure that what we just pushed looks like what the server thinks
	s, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error getting blocked services %w", err)
	}

	slices.Sort(blocked.IDs)
	slices.Sort(s.IDs)
	if !slices.Equal(blocked.IDs, s.IDs) {
		return fmt.Errorf("service lists unequal: %v %v", blocked.IDs, s.IDs)
	}
>>>>>>> Stashed changes

	return nil

}

func init() {
	servicesCmd.AddCommand(serviceUpdateCmd)
	serviceUpdateCmd.Flags().StringSliceVar(&toPermit, "permit", []string{}, "CSV of services to permit")
	serviceUpdateCmd.Flags().StringSliceVar(&toBlock, "block", []string{}, "CSV of services to block")

}
