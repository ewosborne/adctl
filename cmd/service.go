/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/ewosborne/adctl/common"
	"github.com/spf13/cobra"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "service",
	Short: "Alter filtered services",
}

// serviceUpdateCmd represents the update command
var serviceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Block or unblock one or more services",
	RunE:  UpdateServiceCmdE,
}

// serviceListCmd represents the list command
var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blockable or blocked services",
}

// serviceListAllCmd represents the all command
var serviceListAllCmd = &cobra.Command{
	Use:   "all",
	Short: "List all blockable services",
	RunE:  ListAllCmdE,
}

var serviceListBlockedCmd = &cobra.Command{
	Use:   "blocked",
	Short: "List all blocked services",
	RunE:  serviceListBlockedCmdE,
}

func init() {
	rootCmd.AddCommand(servicesCmd)

	servicesCmd.AddCommand(serviceUpdateCmd)
	serviceUpdateCmd.Flags().StringSliceVarP(&toUnblock, "unblock", "u", []string{}, "CSV of services to unblock")
	serviceUpdateCmd.Flags().StringSliceVarP(&toBlock, "block", "b", []string{}, "CSV of services to block")

	servicesCmd.AddCommand(serviceListCmd)

	serviceListCmd.AddCommand(serviceListAllCmd)

	rewriteListCmd.AddCommand(serviceListBlockedCmd)

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
		if svc == "all" {
			return nil, fmt.Errorf("cannot block all services")
		}
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
