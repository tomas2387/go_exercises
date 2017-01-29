package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}


func SqrtFast(x float64)  (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	for i:=1; i < 10; i++ {
		dividend := math.Pow(z, 2) - x
		divisor := float64(2) * z
		z = z - (dividend / divisor)
	}
	return z, nil
}

func SqrtRight(x float64)  (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	delta := float64(1)
	for  delta > 1e-10 {
		dividend := math.Pow(z, 2) - x
		divisor := float64(2) * z
		lastZ := z
		z = z - (dividend / divisor)
		delta = math.Abs(z - lastZ)
	}
	return z, nil
}

func main() {
	fast, err := SqrtFast(2)
	if err != nil {
		fmt.Println(err);
		return
	}
	correct, err := SqrtRight(2)
	if err != nil {
		fmt.Println(err);
		return
	}

	original := math.Sqrt(2)

	fmt.Printf("This is fast %v \nAnd this is correct %v \nBut all must equal %v", fast, correct, original)
}
