package main

import (
	"Assignment2/calc"
	"fmt"
)

func main() {
	fmt.Println("Additon ", calc.Add(10, 20))
	sub := calc.Sub(50, 30)
	fmt.Println("Sub value is", sub)
	multiplication := calc.Mul(10, 30)
	fmt.Println("Multiplication value is", multiplication)
	d := calc.Div(50, 10)
	fmt.Println("division output is", d)

}
