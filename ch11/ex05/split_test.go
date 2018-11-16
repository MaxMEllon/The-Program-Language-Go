package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	suite := []struct {
		sep    string
		data   string
		expect []string
	}{
		{
			":",
			"1:2:3",
			[]string{"1", "2", "3"},
		},
		{
			",",
			"a,b,c",
			[]string{"a", "b", "c"},
		},
		{
			"\u200d",
			"👨‍👩‍👧‍👦",
			[]string{"👨", "👩", "👧", "👦"},
		},
		{
			",",
			"👴🏻",
			[]string{"👴🏻"},
		},
		{
			"🏻",
			"👴🏻",
			[]string{"👴", ""},
		},
		{
			"\u200d",
			"👨‍😾‍👧‍👦",
			[]string{"👨", "😾", "👧", "👦"},
		},
	}

	for _, tc := range suite {
		act := strings.Split(tc.data, tc.sep)
		if len(act) != len(tc.expect) {
			t.Errorf("actual: %q,  expect: %q", act, tc.expect)
		}

		for idx, a := range act {
			expect := tc.expect[idx]
			if expect != a {
				t.Errorf("actual: %q,  expect: %q", act, tc.expect)
			}
		}
	}
}
