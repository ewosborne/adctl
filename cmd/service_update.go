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
var toUnblock []string
var toBlock []string

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
	if len(toBlock) == 0 && len(toUnblock) == 0 {
		return fmt.Errorf("need permit or blocked flag")
	}
	// first tidy up.
	toBlock = unique(toBlock)
	toUnblock = unique(toUnblock)

	svcs := ServiceLists{block: toBlock, permit: toUnblock}

	err := updateServices(svcs)
	if err != nil {
		return fmt.Errorf("error updating services %w", err)
	}

	return nil
}

func computeNewBlocks(currentlyBlocked AllBlockedServices, changes ServiceLists) ([]string, error) {
	ret := []string{}
	svcmap := make(map[string]bool)

	// take currentlyBlocked.IDs and enter them into the map
	//fmt.Println("currently blocked", currentlyBlocked.IDs)
	for _, svc := range currentlyBlocked.IDs {
		svcmap[svc] = true
	}

	// add all changes.block

	debugLogger.Println("to block", changes.block)
	for _, svc := range changes.block {
		svcmap[svc] = true
	}

	// subtract all changes.permit
	debugLogger.Println("to permit", changes.permit)
	for _, svc := range changes.permit {
		svcmap[svc] = false
	}

	// turn back into a list of services which is the new thing to push
	for k := range svcmap {
		if svcmap[k] {
			ret = append(ret, k)
		}
	}

	/// special case to disable all
	for _, k := range changes.permit {
		if k == "all" {
			ret = []string{}
		}
	}

	// clean up.  no dups since it came from map keys.
	slices.Sort(ret)

	debugLogger.Print("final set to enable ", ret)
	// return it

	return ret, nil
}

func updateServices(svcs ServiceLists) error {

	// note that blocked has a schedule as well.  That just gets passed transparently through, I don't touch it.
	blocked, err := GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error calling GetBlockedServices %w", err)
	}

	newList, err := computeNewBlocks(blocked, svcs)
	if err != nil {
		return fmt.Errorf("error computing new blocks: %w", err)
	}

	var requestBody = make(map[string]any)
	requestBody["ids"] = newList
	requestBody["schedule"] = blocked.Schedule

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return err
	}

	baseURL.Path = "/control/blocked_services/update"

	debugLogger.Println("going to update with", requestBody)

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
	if !slices.Equal(newList, s.IDs) {
		return fmt.Errorf("service lists unequal: %v %v", newList, s.IDs)
	}

	// TODO: check to make sure that what we just pushed looks like what the server thinks
	s, err = GetBlockedServices()
	if err != nil {
		return fmt.Errorf("error getting blocked services %w", err)
	}

	slices.Sort(blocked.IDs)
	slices.Sort(s.IDs)
	if !slices.Equal(newList, s.IDs) {
		return fmt.Errorf("service lists unequal: %v %v", newList, s.IDs)
	}

	err = PrintBlockedServices()
	if err != nil {
		return err
	}

	return nil

}

func init() {
	servicesCmd.AddCommand(serviceUpdateCmd)
	serviceUpdateCmd.Flags().StringSliceVarP(&toUnblock, "unblock", "u", []string{}, "CSV of services to unblock")
	serviceUpdateCmd.Flags().StringSliceVarP(&toBlock, "block", "b", []string{}, "CSV of services to block")

}
