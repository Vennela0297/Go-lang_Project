package main

import (
	"fmt"

	"mathcal/addition"
	"mathcal/division"
	"mathcal/multiply"
	"mathcal/subtraction"
)

func main() {

	x := addition.Add(3, 4)
	fmt.Println(x)

	y := subtraction.Sub(4, 3)
	fmt.Println(y)

	z := multiply.Mul(4, 3)
	fmt.Println(z)

	w := division.Div(8, 4)
	fmt.Println(w)

	fmt.println("hello git installed")

}
