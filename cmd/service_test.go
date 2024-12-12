package cmd

import (
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

}

func Test_computeNewBlock(t *testing.T) {
	// try an internal test thing

	// computeNewBlocks(currentlyBlocked AllBlockedServices, changes ServiceLists) ([]string, error)

	var tt = []struct {
		currentlyBlocked []string
		block            []string
		permit           []string
		expected         []string
	}{
		{
			currentlyBlocked: []string{"yy"},
			block:            []string{"4chan"},
			permit:           []string{},
			expected:         []string{"4chan", "yy"},
		},
		{
			currentlyBlocked: []string{},
			block:            []string{"4chan", "yy"},
			permit:           []string{},
			expected:         []string{"4chan", "yy"},
		},
		{
			currentlyBlocked: []string{},
			block:            []string{},
			permit:           []string{},
			expected:         []string{},
		},
		{
			currentlyBlocked: []string{"yy", "4chan"},
			block:            []string{"8chan"},
			permit:           []string{"all"},
			expected:         []string{},
		},
		{
			currentlyBlocked: []string{"yy", "4chan"},
			block:            []string{"8chan"},
			permit:           []string{},
			expected:         []string{"4chan", "8chan", "yy"},
		},
	}

	for _, entry := range tt {
		t.Log(entry)
		cb := AllBlockedServices{IDs: entry.currentlyBlocked}
		changes := ServiceLists{block: entry.block, permit: entry.permit}

		res, err := computeNewBlocks(cb, changes)
		if err != nil {
			t.Errorf("error in test %s", err)
		}
		t.Log("expected", entry.expected, "got", res)
		if slices.Compare(entry.expected, res) != 0 {
			t.Errorf("compared wrong expected:%v, res %v, %v", entry.expected, res, slices.Compare(entry.expected, res))
		}
	}
}
