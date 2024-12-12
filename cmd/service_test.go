package cmd

import (
	"fmt"
	"slices"
	"testing"
)

func Test_ListAll(t *testing.T) {
	_, err := GetAllServices()
	if err != nil {
		t.Error(err)
	}

}
func Test_ListBlocked(t *testing.T) {
	_, err := GetBlockedServices()
	if err != nil {
		t.Error(err)
	}

}

func Test_Update(t *testing.T) {
	// TODO my APIs are weird, I'm not sure this is the right approach, skipping for now.

	t.Skip()
	// get old test list
	original_services, _ := GetBlockedServices()
	fmt.Println("original services slice", original_services.IDs)
	// block a service
	updateServices([]string{}, []string{"zhihu"})
	// get new block list
	post_change, _ := GetBlockedServices()
	fmt.Println("old", original_services.IDs, "new", post_change.IDs)

	// make sure it's accurate
	// ....
	// revert to what it was
	// ...
	updateServices(original_services.IDs, []string{})
	// make sure the revert tok
	reverted, _ := GetBlockedServices()

	slices.Sort(original_services.IDs)
	slices.Sort(reverted.IDs)
	if slices.Compare(reverted.IDs, post_change.IDs) != 0 {
		t.Errorf("service lists not the same: %v %v", original_services.IDs, reverted.IDs)
	}

}
