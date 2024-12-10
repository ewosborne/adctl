/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	*/

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
