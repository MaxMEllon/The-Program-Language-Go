package popcount

import (
	"fmt"
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch09/ex02/popcount"
)

func TestPopCount(t *testing.T) {
	testCase := []struct {
		data     uint64
		expected int
	}{
		{
			data:     0x11111111111,
			expected: 11,
		},
		{
			data:     0x1111,
			expected: 4,
		},
		{
			data:     0x1,
			expected: 1,
		},
		{
			data:     0x11,
			expected: 2,
		},
		{
			data:     0x111111111111,
			expected: 12,
		},
	}
	for _, tc := range testCase {
		count := popcount.PopCount(tc.data
		)
		if count != tc.expected {
			panic(fmt.Sprintf("count is %d, want %d\n", count, tc.expected))
		}
	}
}
