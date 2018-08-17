package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	for _, test := range []struct {
		data   string
		expect string
	}{
		{"", ""},
		{"1", "1"},
		{"-1", "-1"},
		{"1000.000", "1,000.000"},
		{"-1000.000", "-1,000.000"},
	} {
		result := comma(test.data)
		if result != test.expect {
			t.Errorf("result = %q, but want %q", result, test.expect)
		}
	}
}
