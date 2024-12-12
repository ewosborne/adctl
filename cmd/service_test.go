package cmd

import "testing"

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

	// TODO: also need to make sure I'm passing a schedule through transparently.

	t.Skip()
	// get old blocked list

	// original_blocked, _ := GetBlockedServices()

	// // block one new service, 'iqiyi'
	// new_blocked := AllBlockedServices{}
	// new_blocked.IDs := append(original_blocked.IDs, "iqiyi")

	// // push the block
	// // hrmm todo what should this look like?  goes back to whether the updateServices() api is right.

	// make sure what comes back is what we sent
	// revert to original blocked list

}
