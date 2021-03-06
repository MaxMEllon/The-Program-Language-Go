package params

import "testing"

func TestPack(t *testing.T) {
	suite := []struct {
		st       interface{}
		expected string
	}{
		{
			st: &struct {
				foo   string            `http:"foo"`
				bar   int               `http:"bar"`
				hoge  bool              `http:"hoge"`
				poge  []int             `http:"poge"`
				ababa map[string]string `http:"ababa"`
			}{
				foo:  "foo",
				bar:  2,
				hoge: true,
				poge: []int{1, 2, 3},
				ababa: map[string]string{
					"foo": "bar",
				},
			},
			expected: "ababa.foo=bar&bar=2&foo=foo&hoge=true&poge%5B%5D=1&poge%5B%5D=2&poge%5B%5D=3",
		},
		{
			st: &struct {
				m map[bool]bool `http:"m"`
			}{
				m: map[bool]bool{
					true: true,
				},
			},
			expected: "m.true=true",
		},
		{
			st: &struct {
				arr [][]int `http:"arr"`
			}{
				arr: [][]int{{1, 2, 3}, {2, 3, 4}},
			},
			expected: "arr%5B%5D=1%2C2%2C3&arr%5B%5D=2%2C3%2C4",
		},
	}
	for _, c := range suite {
		if actual, err := Pack(c.st); actual != c.expected || err != nil {
			if err == nil {
				t.Errorf("expected: %v, but got: %v", c.expected, actual)
			} else {
				t.Errorf("unexpected error %v\n", err)
			}
		}
	}
}
