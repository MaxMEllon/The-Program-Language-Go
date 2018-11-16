package intset

import (
	"math/rand"
	"testing"
	"time"
)

var seed int64

func init() {
	seed = time.Now().UTC().UnixNano()
}

func BenchmarkIntSetAdd(b *testing.B) {
	rng := rand.New(rand.NewSource(seed))
	s := &IntSet{}
	for i := 0; i < b.N; i++ {
		s.Add(rng.Intn(0xffffffff))
	}
}
