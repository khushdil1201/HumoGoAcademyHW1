package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter number:")

	fmt.Scan(&number)

	fmt.Println("Result: ", number/100)
}
