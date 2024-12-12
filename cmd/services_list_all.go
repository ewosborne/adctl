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

type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type AllServices struct {
	AllServices []Service `json:"blocked_services"`
}

type ServiceMap struct {
	ID2Name map[string]string
	Name2ID map[string]string
}

func NewServiceMap() ServiceMap {
	return ServiceMap{
		ID2Name: make(map[string]string),
		Name2ID: make(map[string]string),
	}
}

// serviceListAllCmd represents the all command
var serviceListAllCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	RunE:  ListAllCmdE,
}

func ListAllCmdE(cmd *cobra.Command, args []string) error {

	err := PrintAllServices()
	if err != nil {
		return fmt.Errorf("error somewhere %w", err)
	}

	return nil
}

// TODO: make this json or text
func PrintAllServices() error {

	smap, err := GetAllServices()
	name2id := smap.Name2ID

	if err != nil {
		return err
	}

	// print name2id
	// if json
	b, err := json.MarshalIndent(name2id, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	// if text

	// s := slices.Collect(maps.Keys(name2id))
	// sort.Strings(s)

	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	// defer w.Flush() // TODO I can't decide if this is dumb or not
	// fmt.Fprintf(w, "Name\tID\n")
	// fmt.Fprintf(w, "====\t==\n")

	// for _, k := range s {
	// 	fmt.Fprintf(w, "%s\t%s\n", k, name2id[k])
	// }
	return nil
}

func GetAllServices() (ServiceMap, error) {

	ret := NewServiceMap()

	id2name := ret.ID2Name
	name2id := ret.Name2ID

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

	// TODO: marshal body into something that pulls out name and ID.  AllServices{ Service } however I do that.

	// this is a very confusing mess of nested structs

	var s AllServices
	json.Unmarshal(body, &s)

	for _, x := range s.AllServices {
		id2name[x.ID] = x.Name
		name2id[x.Name] = x.ID

	}

	return ret, nil
}

func init() {
	listCmd.AddCommand(serviceListAllCmd)
}
