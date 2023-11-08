package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b float64

	fmt.Println("Enter a:")

	fmt.Scan(&a)

	fmt.Println("Enter b:")

	fmt.Scan(&b)

	if math.IsNaN(math.Sqrt(a * b)) {
       fmt.Println("Error! a*b>=0 !")
	}else{
		fmt.Println("Result:",math.Sqrt(a*b))

	}
	


	
}
