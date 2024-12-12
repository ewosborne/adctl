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

// serviceListBlockedCmd represents the blocked command
var serviceListBlockedCmd = &cobra.Command{
	Use:   "blocked",
	Short: "A brief description of your command",
	RunE:  serviceListBlockedCmdE,
}

type AllBlockedServices struct {
	Schedule map[string]any `json:"schedule"`
	IDs      []string       `json:"ids"`
}

func serviceListBlockedCmdE(cmd *cobra.Command, args []string) error {

	err := PrintBlockedServices()
	if err != nil {
		return fmt.Errorf("error somewhere %w", err)
	}

	return nil
}

type BlockedWithCount struct {
	Count int      `json:"count"`
	IDs   []string `json:"IDs"`
}

func PrintBlockedServices() error {

	s, err := GetBlockedServices()

	if err != nil {
		return err
	}

	// json
	var x BlockedWithCount
	x.Count = len(s.IDs)
	x.IDs = s.IDs
	//x := BlockedWithCount{Count: len(s.IDs), AllBlockedServices.AllBlockedServices: s.IDs}
	b, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	// text
	// if len(s.IDs) == 0 {
	// 	fmt.Println("no services blocked")
	// } else {
	// 	allServices, err := GetAllServices()
	// 	if err != nil {
	// 		return fmt.Errorf("error getting all services: %w", err)
	// 	}
	// 	for _, x := range s.IDs {
	// 		fmt.Println("svc blocked", allServices.ID2Name[x])
	// 	}
	// }

	return nil
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

	var s AllBlockedServices
	json.Unmarshal(body, &s)

	return s, nil
}

func init() {
	listCmd.AddCommand(serviceListBlockedCmd)

}
