package main

import (
	"calci/operations/add"
	"calci/operations/div"
	"calci/operations/mul"
	"calci/operations/sub"
	"fmt"
	"reflect"
)

func typeCheck(i int) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		return true
	default:
		fmt.Println(reflect.TypeOf(i))
		return false
	}
}

func main() {

	fmt.Println("Please enter the value for X :")
	var x, y int
	fmt.Scanln(&x)
	//fmt.Print(x)
	if !typeCheck(x) {
		fmt.Println("Please enter the value for X :")
		fmt.Scanln(&x)
	}
	fmt.Println("Please enter the value for Y :")
	fmt.Scanln(&y)
	if !typeCheck(y) {
		fmt.Println("Please enter the value for Y :")
		fmt.Scanln(&y)
	}

	fmt.Println("Please select any one of the operations to be performed\n1. Addition\n2.Subtraction\n3.Multiplication\n4.Division")
	var sw int
	fmt.Scanln(&sw)

	switch sw {
	case 1:
		add.Add(x, y)
	case 2:
		sub.Sub(x, y)
	case 3:
		mul.Mul(x, y)
	case 4:
		div.Div(x, y)
	default:
		fmt.Println("Your selected operation is not available, Thank you")
	}
}
