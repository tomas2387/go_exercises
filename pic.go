package main

import "golang.org/x/tour/pic"

func Pic1(dx, dy int) [][]uint8 {
	sliceDy := make([][]uint8, dy)
	for x := 0; x < dy; x++ {
		sliceDy[x] = make([]uint8, dx)

		for y := 0; y < dx; y++ {
			sliceDy[x][y] = uint8((x + y) / 2);
		}
	}
	return sliceDy
}

func Pic2(dx, dy int) [][]uint8 {
	sliceDy := make([][]uint8, dy)
	for x := 0; x < dy; x++ {
		sliceDy[x] = make([]uint8, dx)

		for y := 0; y < dx; y++ {
			sliceDy[x][y] = uint8(x * y);
		}
	}
	return sliceDy
}

func Pic3(dx, dy int) [][]uint8 {
	sliceDy := make([][]uint8, dy)
	for x := 0; x < dy; x++ {
		sliceDy[x] = make([]uint8, dx)

		for y := 0; y < dx; y++ {
			sliceDy[x][y] = uint8(x ^ y);
		}
	}
	return sliceDy
}

func Pic4(dx, dy int) [][]uint8 {
	sliceDy := make([][]uint8, dy)
	for x := 0; x < dy; x++ {
		sliceDy[x] = make([]uint8, dx)

		for y := 0; y < dx; y++ {
			div := uint8(x)
			if y > 0 {
				div = uint8(2 * (x / y) ^ x)
			}
			sliceDy[x][y] = div
		}
	}
	return sliceDy
}

func main() {
	pic.Show(Pic3)
}
