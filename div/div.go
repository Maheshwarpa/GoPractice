package div

import "fmt"

func Div(i, j int) (int, int) {
	var q int
	var r int

	q = i / j
	r = i % j

	fmt.Printf("Division of %d and %d for quotient %d and remainder is %d", i, j, q, r)

	return q, r
}
