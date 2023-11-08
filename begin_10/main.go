package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Enter a:")

	fmt.Scan(&a)

	fmt.Println("Enter b:")

	fmt.Scan(&b)

	if a*b == 0 {
		fmt.Println("Error! a*b!=0")
	} else {
		fmt.Println("Sum:", math.Pow(a, 2)+math.Pow(b, 2))
		fmt.Println("Multiplication:", math.Pow(a, 2)*math.Pow(b, 2))
		fmt.Println("Subtraction:", math.Pow(a, 2)-math.Pow(b, 2))
		fmt.Println("Division:", math.Pow(a, 2)/math.Pow(b, 2))

	}

}
