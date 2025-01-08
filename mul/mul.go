package mul

import "fmt"

func Mul(i, j int) int {
	fmt.Printf("Multiplication of %d and %d is: %d", i, j, (i * j))
	return i * j
}
