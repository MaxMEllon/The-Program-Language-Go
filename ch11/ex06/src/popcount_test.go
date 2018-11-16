package popcount

import (
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch02/ex04/src/popcount"
)

var blank int

func BenchmarkIsPopCountByVerbose(b *testing.B) {
	blank = 0
	for i := 0; i < b.N; i++ {
		blank = popcount.PopCountByVerbose(0xffffffff)
		blank--
	}
	b.Log(blank)
}

func BenchmarkIsPopCountByFor(b *testing.B) {
	blank = 0
	for i := 0; i < b.N; i++ {
		blank = popcount.PopCountByFor(0xffffffff)
		blank--
	}
	b.Log(blank)
}

func BenchmarkIsPopCount(b *testing.B) {
	blank = 0
	for i := 0; i < b.N; i++ {
		blank = popcount.PopCount(0xffffffff)
		blank--
	}
	b.Log(blank)
}
