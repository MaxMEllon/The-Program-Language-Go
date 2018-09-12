package intset

import (
	"bytes"
	"fmt"
)

const SIZE = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func NewIntSet(xs ...int) *IntSet {
	s := &IntSet{}
	s.AddAll(xs...)
	return s
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/SIZE, uint(x%SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Elems() []int {
	elems := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, SIZE*i+j)
			}
		}
	}
	return elems
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func popcount(x uint) int {
	var count int
	for i := 0; i < SIZE; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	s.words[word] |= 1 << bit
	s.words[word] ^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	new_s := &IntSet{}
	new_s.words = make([]uint, len(s.words))
	copy(new_s.words, s.words)
	return new_s
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
