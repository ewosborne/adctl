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

	// get list of what's currently blocked, store it
	//  add some new services, subtract anything blocked (?).  maybe random among some set of ones I don't use?
	// push that list, fetch it, make sure it matches.
	// put the original list back.

}
