package popcount

import (
	"fmt"
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch09/ex02/popcount"
)

func TestPopCount(t *testing.T) {
	count := popcount.PopCount(0x11111111111)
	if count != 20 {
		panic(fmt.Sprintf("count is %d, want 20\n", count))
	}
}