package tempconv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCToK(t *testing.T) {
	k := CToK(AbsoluteZeroC)
	assert.Equal(t, k, Kelvin(0), "摂氏 絶対零度は ケルビン0度である")
}

func TestKToC(t *testing.T) {
	c := KToC(Kelvin(float64(0)))
	assert.Equal(t, c, AbsoluteZeroC, "ケルビン0度は 絶対零度である")
}
