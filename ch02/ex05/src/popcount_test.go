package popcount

import (
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch02/ex05/src/popcount"
)

const (
	target = 0xfffffffffffff
)

func BenchmarkIsPopCountByVerbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByVerbose(target)
	}
}

func BenchmarkIsPopCountByFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByFor(target)
	}
}

func BenchmarkIsPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(target)
	}
}

func BenchmarkIsPopCountByTheLowestBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByTheLowestBitClear(target)
	}
}
