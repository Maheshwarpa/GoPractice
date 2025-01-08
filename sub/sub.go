package sub

import "fmt"

func Sub(i, j int) int {
	fmt.Printf("Subtraction of %d and %d is: %d", i, j, (i - j))
	return i - j
}
