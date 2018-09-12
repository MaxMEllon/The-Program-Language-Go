package intset

import (
	"reflect"
	"testing"
)

func Test_IntSet_Elems(t *testing.T) {
	suite := []struct {
		s        *IntSet
		expected []int
	}{
		{
			NewIntSet(1, 2, 3, 4, 5, 6, 7, 8, 9),
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			NewIntSet(1, 2),
			[]int{1, 2},
		},
		{
			NewIntSet(),
			[]int{},
		},
		{
			NewIntSet(0xffffffff),
			[]int{},
		},
	}

	for _, testCase := range suite {
		actual := testCase.s.Elems()
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("actual %v want %v", actual, testCase.expected)
		}
	}
}
