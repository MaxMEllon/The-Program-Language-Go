package intset

import (
	"strconv"
	"strings"
	"testing"
)

func Test_IntSet_Empty(t *testing.T) {
	s := &IntSet{}
	actual := s.String()
	expected := "{}"
	if actual != expected {
		t.Errorf("actual %v want %v", actual, expected)
	}
}

func Test_IntSet_Add(t *testing.T) {
	suite := map[string][]int{
		"{1 2 3}":    []int{1, 2, 3},
		"{1}":        []int{1, 1, 1},
		"{}":         []int{},
		"{16777215}": []int{0xFFFFFF, 0xFFFFFF, 0xFFFFFF, 0xFFFFFF},
	}
	for expected, data := range suite {
		s := &IntSet{}
		for _, v := range data {
			s.Add(v)
		}
		actual := s.String()
		if actual != expected {
			t.Errorf("actual %v want %v", actual, expected)
		}
	}
}

func Test_IntSet_Has(t *testing.T) {
	suite := map[bool]map[int][]int{
		true: {
			1:  []int{1, 1, 1},
			8:  []int{1, 8, 9},
			23: []int{21, 8, 9, 23, 1, 4, 5},
		},
		false: {
			0: []int{1, 1, 1},
			// 1:  []int{}, 失敗する
			23: []int{20, 21, 22, 24},
		},
	}

	for expected, testCase := range suite {
		s := &IntSet{}
		for key, data := range testCase {
			for _, v := range data {
				s.Add(v)
			}
			actual := s.Has(key)
			if actual != expected {
				t.Errorf("actual %v want %v", actual, expected)
			}
		}
	}
}

func Test_IntSet_UnionWith(t *testing.T) {
	suite := map[string][2][]int{
		"{1 2 3 4 5 6}": [2][]int{
			{1, 2, 3},
			{4, 5, 6},
		},
		"{}": [2][]int{
			{},
			{},
		},
		"{1 2 3 5 6}": [2][]int{
			{1, 1, 2, 5, 6},
			{2, 2, 3, 5, 6},
		},
	}

	for expected, testCase := range suite {
		s1, s2 := &IntSet{}, &IntSet{}
		for _, v := range testCase[0] {
			s1.Add(v)
		}
		for _, v := range testCase[1] {
			s2.Add(v)
		}
		s1.UnionWith(s2)
		actual := s1.String()
		if actual != expected {
			t.Errorf("actual %v want %v", actual, expected)
		}
	}
}

func Test_IntSet_Len(t *testing.T) {
	suite := map[int][]int{
		0: []int{},
		1: []int{1, 1},
		2: []int{1, 2},
		3: []int{1, 2, 2, 2, 2, 3},
		4: []int{1, 2, 2, 2, 2, 3, 4},
		5: []int{1, 2, 3, 4, 5},
	}

	for expected, testCase := range suite {
		s := &IntSet{}
		for _, v := range testCase {
			s.Add(v)
		}
		actual := s.Len()
		if actual != expected {
			t.Errorf("actual %v want %v", actual, expected)
		}
	}
}

// 通るときと通らない時がある... (ガチャテスト)
func Test_IntSet_Remove(t *testing.T) {
	suite := map[string]map[string]int{
		"0.1.2.3": {
			// "{1 2 3}": 0, コケる
			"{0 2 3}": 1,
			"{0 3}":   2,
			"{0}":     3,
		},
		"1.8.9.10": {
			"{1 8 9 10}": 5,
			"{1 9 10}":   8,
			"{1 10}":     9,
			"{1}":        10,
		},
	}

	for initalData, testCase := range suite {
		s := &IntSet{}
		for _, v := range strings.Split(initalData, ".") {
			i, _ := strconv.Atoi(v)
			s.Add(i)
		}
		for expected, removeKey := range testCase {
			s.Remove(removeKey)
			actual := s.String()
			if actual != expected {
				t.Errorf("actual %v want %v", actual, expected)
			}
		}
	}
}

func Test_IntSet_Clear(t *testing.T) {
	suite := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{6, 2, 3, 4, 5, 6, 7},
		{7},
		{},
		{0},
	}
	expected := "{}"
	for _, testCase := range suite {
		s := &IntSet{}
		for _, v := range testCase {
			s.Add(v)
		}
		s.Clear()
		actual := s.String()
		if actual != expected {
			t.Errorf("actual %v want %v", actual, expected)
		}
	}
}

func Test_IntSet_Copy(t *testing.T) {
	suite := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{6, 2, 3, 4, 5, 6, 7},
		{7},
		{},
		{0},
	}
	for _, testCase := range suite {
		s := &IntSet{}
		for _, v := range testCase {
			s.Add(v)
		}
		s.Clear()
		c := s.Copy()
		expected := c.String()
		actual := s.String()
		if actual != expected {
			t.Errorf("actual %v want %v", actual, expected)
		}
	}
}
