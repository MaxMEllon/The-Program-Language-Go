package weightconv

import (
	"fmt"
)

// Gram グラム
type Gram float64

// Pound ポンド
type Pound float64

func (g Gram) String() string {
	return fmt.Sprintf("%g g", g)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

const (
	f = 0.0022046
)

// GToP Gram to Pound
func GToP(g Gram) Pound {
	return Pound(float64(g) * f)
}

// PToG Pound to Gram
func PToG(p Pound) Gram {
	return Gram(float64(p) / f)
}
