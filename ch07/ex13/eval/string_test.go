package eval

import (
	"testing"
)

func Test_Eval_String(t *testing.T) {
	suite := []struct {
		data     string
		expected string
	}{
		{
			data:     "pow(x, 3) + pow(y, 3)",
			expected: "pow(x, 3.000000) + pow(y, 3.000000)",
		},
		{
			data:     "sin(x,3) + pow(y, 3)",
			expected: "sin(x, 3.000000) + pow(y, 3.000000)",
		},
		{
			data:     "5/9*(F-32)",
			expected: "5.000000 / 9.000000 * F - 32.000000",
		},
	}

	for _, testCase := range suite {
		var actual string
		if expr, err := Parse(testCase.data); err == nil {
			actual = expr.String()
		} else {
			t.Errorf("panic: %v", err)
		}
		if actual != testCase.expected {
			t.Errorf("actual %v want %v", actual, testCase.expected)
		}
	}
}
