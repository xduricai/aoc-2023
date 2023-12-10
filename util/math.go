package util

import "math"

// I stole this from here ^_^: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// expects the discriminant to be non-negative
func SolveQuadraticEquation(a float64, b float64, c float64) (float64, float64) {
	d := math.Pow(float64(b), 2) - (4 * a * c)

	x1 := (-b - math.Sqrt(d)) / (2 * a)
	x2 := (-b + math.Sqrt(d)) / (2 * a)

	return x1, x2
}
