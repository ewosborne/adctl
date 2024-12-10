/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// listBlockedCmd represents the listAll command
var listBlockedCmd = &cobra.Command{
	Use:   "listBlocked",
	Short: "List all services AGH has blocked",
	RunE:  ListBlockedCmdE,
}

// TODO: do this one first since it's importantest.

func ListBlockedCmdE(cmd *cobra.Command, args []string) error {

	err := PrintBlockedServices()
	if err != nil {
		return fmt.Errorf("error somewhere %w", err)
	}

	return nil
}

type AllBlockedServices struct {
	IDs []string `json:"ids"`
}

func GetBlockedServices() (AllBlockedServices, error) {

	// get the data

	ret := AllBlockedServices{}

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}
	baseURL.Path = "/control/blocked_services/get"

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	// TODO: marshal body into something that pulls out name and ID.  AllServices{ Service } however I do that.

	// this is a very confusing mess of nested structs

	var s AllBlockedServices
	json.Unmarshal(body, &s)

	return s, nil
}

func PrintBlockedServices() error {

	s, err := GetBlockedServices()

	if err != nil {
		return err
	}

	if len(s.IDs) == 0 {
		fmt.Println("no services blocked")
	} else {
		allServices, err := GetAllServices()
		if err != nil {
			return fmt.Errorf("error getting all services: %w", err)
		}
		for _, x := range s.IDs {
			// TODO return service name, not id.
			fmt.Println("svc blocked", allServices.ID2Name[x])
		}
	}
	return nil
}

func init() {
	serviceCmd.AddCommand(listBlockedCmd)
}
