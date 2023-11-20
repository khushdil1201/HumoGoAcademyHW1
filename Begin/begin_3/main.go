package main

import "fmt"

func main() {
	var a, b float32

	fmt.Println("Enter a:")

	fmt.Scan(&a)

	fmt.Println("Enter b:")

	fmt.Scan(&b)

	fmt.Println("Result:", 2*(a+b))

}
