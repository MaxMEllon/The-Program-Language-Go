package tempconv

import (
	"fmt"
)

// Celsius 摂氏
type Celsius float64

// Fahrenheit 華氏
type Fahrenheit float64

// Kelvin 絶対
type Kelvin float64

const (
	// AbsoluteZeroC 絶対零度
	AbsoluteZeroC Celsius = -273.15
	// FreezingC 凍る温度
	FreezingC Celsius = 0
	// BoilingC 沸騰温度
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g °K", k)
}

// CToF Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(float64(c) - float64(AbsoluteZeroC))
}

// FToC Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK Fahrenheit to Celsius
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

// KToC Fahrenheit to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(float64(k) + float64(AbsoluteZeroC))
}

// KToF Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
