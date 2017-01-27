package main

import (
	"fmt"
	"math"
)

func SqrtFast(x float64) float64 {
	z := float64(1)
	for i:=1; i < 10; i++ {
		dividend := math.Pow(z, 2) - x
		divisor := float64(2) * z
		z = z - (dividend / divisor)
	}
	return z
}

func SqrtRight(x float64) float64 {
	z := float64(1)
	delta := float64(1)
	for  delta > 1e-10 {
		dividend := math.Pow(z, 2) - x
		divisor := float64(2) * z
		lastZ := z
		z = z - (dividend / divisor)
		delta = math.Abs(z - lastZ)
	}
	return z
}

func main() {
	fmt.Printf("This is fast %v \nAnd this is correct %v \nBut all must equal %v", SqrtFast(2), SqrtRight(2), math.Sqrt(2))
}
