package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func findImaginaryQuadraticRoots(a, b, c complex128) (complex128, complex128) {
	negB := -1 * b
	bSqrd := b * b

	stepOne := bSqrd - (4 * a * c)
	stepTwo := cmplx.Sqrt(stepOne)
	stepThree := negB - stepTwo
	stepFour := negB + stepTwo

	ansOne := stepThree / (2 * a)
	ansTwo := stepFour / (2 * a)

	return ansOne, ansTwo
}

func findQuadraticRoots(a, b, c float64) (float64, float64) {
	negB := -1 * b
	bSqrd := b * b

	stepOne := bSqrd - (4 * a * c)

	if stepOne < 0 {
		fmt.Println("There are two imaginary roots:")
	} else if stepOne == 0 {
		fmt.Println("There is one real root:")
	} else {
		fmt.Println("There are two real roots:")
	}

	stepTwo := math.Sqrt(stepOne)
	stepThree := negB - stepTwo
	stepFour := negB + stepTwo

	ansOne := stepThree / (2 * a)
	ansTwo := stepFour / (2 * a)

	return ansOne, ansTwo
}

func main() {
	// 2x^2 + 10x + 8
	x1, x2 := findQuadraticRoots(2, 10, 8) // -4, -1
	fmt.Printf("Roots: %f, %f \n", x1, x2)

	// x^2 + 6x + 9
	x1, x2 = findQuadraticRoots(1, 6, 9) // -3
	fmt.Printf("Roots: %f, %f \n", x1, x2)

	// 64x^2 + 24x - 10
	x1, x2 = findQuadraticRoots(64, 24, -10) // 1/4, -5/8
	fmt.Printf("Roots: %f, %f \n", x1, x2)

	// 2x^2 - 7x - 4
	x1, x2 = findQuadraticRoots(2, -7, -4) // 4, -1/2
	fmt.Printf("Roots: %f, %f \n", x1, x2)

	// x^2 + x + 1
	x1, x2 = findQuadraticRoots(1, 1, 1) // NaN
	fmt.Printf("Roots: %f, %f \n", x1, x2)
	y1, y2 := findImaginaryQuadraticRoots(1, 1, 1) // -1/2 - i(sqrt(3))/2, -1/2 + i(sqrt(3))/2
	fmt.Printf("Roots: %f, %f \n", y1, y2)

	// x^2 + 1
	x1, x2 = findQuadraticRoots(1, 0, 1) // NaN
	fmt.Printf("Roots: %f, %f \n", x1, x2)
	y1, y2 = findImaginaryQuadraticRoots(1, 0, 1) // i, -i
	fmt.Printf("Roots: %f, %f \n", y1, y2)
}
