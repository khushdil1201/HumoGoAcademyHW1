package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter number:")

	fmt.Scan(&number)

	fmt.Println("left side:", number/10)

	fmt.Println("right side:", number%10)

}
