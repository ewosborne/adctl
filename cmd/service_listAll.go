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
	ID   string `json:"id"`
	Name string `json:"name"`
}
type AllServices struct {
	AllServices []Service `json:"blocked_services"`
}

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "List all services AGH knows about",
	RunE:  ListAllCmdE,
}

// TODO: do this one first since it's importantest.

func ListAllCmdE(cmd *cobra.Command, args []string) error {

	err := PrintAllServices()
	if err != nil {
		return fmt.Errorf("error somewhere %w", err)
	}

	return nil
}

func GetAllServices() (map[string]string, map[string]string, error) {

	var id2name = make(map[string]string)
	var name2id = make(map[string]string)

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
		return id2name, name2id, err
	}
	baseURL.Path = "/control/blocked_services/all"

	statusQuery := common.CommandArgs{
		Method: "GET",
		URL:    baseURL,
	}

	body, err := common.SendCommand(statusQuery)
	if err != nil {
		return id2name, name2id, err
	}

	// TODO: marshal body into something that pulls out name and ID.  AllServices{ Service } however I do that.

	// this is a very confusing mess of nested structs

	var s AllServices
	json.Unmarshal(body, &s)

	for _, x := range s.AllServices {
		//fmt.Printf("ID: %s, Name: %s\n", x.ID, x.Name)
		id2name[x.ID] = x.Name
		name2id[x.Name] = x.ID

	}
	//fmt.Printf("%+v\n", s)

	return id2name, name2id, nil
}

func PrintAllServices() error {

	fmt.Print("in PrintAllServices")
	_, name2id, err := GetAllServices()

	if err != nil {
		return err
	}

	for k, v := range name2id {
		fmt.Println("Name:", k, "ID:", v)
	}

	return nil
}

func init() {
	serviceCmd.AddCommand(listAllCmd)
}
