/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// TODO: handle both all and blocked
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List services, all or blocked",

	RunE: ServiceListCmdE,
}

//  blocked services
// {
// 	"ids": [
// 	  "string"
// 	]
//   }

type AllServices struct {
	Services []Service `json:"blocked_services"`
}

type Service struct {
	//IconSVG string   `json:"icon_svg"`
	ID   string `json:"id"`
	Name string `json:"name"`
	//Rules   []string `json:"rules"`
}

func ServiceListCmdE(cmd *cobra.Command, args []string) error {

	// maybe here's where we handle --all or --blocked flags?

	//fmt.Println("all flag", all_services)
	//fmt.Println("blocked flag", blocked_services)

	var flagStr string

	// not sure this is stricly necessary but it doesn't hurt
	if all_services && !blocked_services {
		flagStr = "all"
	} else if blocked_services && !all_services {
		flagStr = "blocked"
	} else {
		return fmt.Errorf("something is wrong with service flags")
	}

	//fmt.Println("service flags", flagStr)
	body, err := GetServiceList(flagStr)

	_ = body
	if err != nil {
		return err
	}

	return nil
}

func GetServiceList(kind string) (bytes.Buffer, error) {

	var ret bytes.Buffer

	// /control/blocked_services/all

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}

	/*
		TODO:
		first call GetBlockedServiceList to get IDs of blocked
		then call GetAllServiceList to get k/v map[string]string  ID to Name, if there's anything blocked.
		then show names of blocked (maybe just []string)

		might want to be able to toggle service blocks on and off via cli, that'd be cool.
	*/

	switch kind {
	case "all":
		baseURL.Path = "/control/blocked_services/all"
	case "blocked":
		baseURL.Path = "/control/blocked_services/get"
	}

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	//fmt.Println(string(body))

	/* TODO: ok here's some design thoughts

	calling the all endpoint gets K/V of all services, ID and Name
	calling the blocked ones get ID of blocked ones

	maybe call the all endpoint and get everything just to get ID/name mapping
		which I want as a map[string]string
	then call blocked to get blocked IDs
	then show blocked names?

	also need to allow for toggle later on.

	*/

	switch kind {
	case "all":
		var allServices AllServices
		if err := json.Unmarshal([]byte(body), &allServices); err != nil {
			log.Fatal(err)
		}

		for _, service := range allServices.Services {
			fmt.Printf("%+v\n", service)
		}
	case "blocked":
		fmt.Println("don't do much with blocked yet")
		fmt.Println(string(body))
	}

	return ret, nil

}

var all_services bool
var blocked_services bool

func init() {
	serviceCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&all_services, "all", "", false, "List all services")
	listCmd.Flags().BoolVarP(&blocked_services, "blocked", "", false, "List blocked services")
	listCmd.MarkFlagsMutuallyExclusive("all", "blocked")

}
