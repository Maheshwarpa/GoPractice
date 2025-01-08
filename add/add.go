package add

import "fmt"

//import "fmt"

func Add(i int, j int) int {
	fmt.Printf("Addition of %d and %d is: %d", i, j, (i + j))
	return i + j
}
