package popcount

import (
	"testing"

	"github.com/maxmellon/The-Program-Language-Go/ch02/ex04/src/popcount"
)

func BenchmarkIsPopCountByVerbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByVerbose(0xffffffff)
	}
}

func BenchmarkIsPopCountByFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByFor(0xffffffff)
	}
}

func BenchmarkIsPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0xffffffff)
	}
}
