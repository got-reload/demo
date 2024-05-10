package example2

import (
	"fmt"
	"math"
)

func Sin(f float64) float64 {
	fmt.Printf("example2.Sin: f: %0.3f\n", f)
	return math.Sin(f)
}

func ShowS(s string) {
	fmt.Printf("example2.ShowS: %s\n", s)
}
