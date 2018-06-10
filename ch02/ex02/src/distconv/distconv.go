package distconv

import (
	"fmt"
)

// Metre メートル
type Metre float64

// Yard ヤード
type Yard float64

func (m Metre) String() string {
	return fmt.Sprintf("%g m", m)
}

func (y Yard) String() string {
	return fmt.Sprintf("%g yd", y)
}

const (
	f = 1.0936
)

// MToY Metre to Yard
func MToY(m Metre) Yard {
	return Yard(float64(m) * f)
}

// YToM Yard to Metre
func YToM(y Yard) Metre {
	return Metre(float64(y) / f)
}
