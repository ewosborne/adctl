/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

/*
{
  "blocked_services": [
    {
      "icon_svg": "string",
      "id": "string",
      "name": "string",
      "rules": [
        "string"
      ]
    }
  ]
}
*/

type Service struct {
	ID   string
	Name string
}
type AllServices struct {
	AllServices []Service
}

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "List all services AGH knows about",
	RunE:  ListAllCmdE,
}

// TODO: do this one first since it's importantest.

func ListAllCmdE(cmd *cobra.Command, args []string) error {

	ret, err := GetAllServices()
	if err != nil {
		return fmt.Errorf("error calling GetAllServices: %w", err)
	}

	PrintAllServices(ret)

	return nil
}

func GetAllServices() (map[string]string, error) {
	var ret map[string]string

	/*
		TODO: get services, populate map with k=ID, v=Name
		and also maybe Name:ID ?  or two maps?  let's try one.

		endpoint is /control/blocked_services/all

		* get the data
		* marshal it into a map, I guess.  map[string]any
		* walk that thing and pull out what I want into a more structured setup?

	*/

	// get the data

	baseURL, err := common.GetBaseURL()
	if err != nil {
		return ret, err
	}
	baseURL.Path = "/control/blocked_services/all"

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return ret, err
	}

	// TODO: marshal body into something that pulls out name and ID

	fmt.Println(string(body))
	return ret, nil
}

func PrintAllServices(data map[string]string) error {

	fmt.Print("in PrintAllServices")
	fmt.Print(data)
	return nil
}

func init() {
	serviceCmd.AddCommand(listAllCmd)
}
