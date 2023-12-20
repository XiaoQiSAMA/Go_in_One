package method

import "math"

func calcTriangle(a, b int) int {
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
