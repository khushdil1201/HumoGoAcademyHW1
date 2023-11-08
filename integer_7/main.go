package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter number:")

	fmt.Scan(&number)

	fmt.Println("sum:", number/10+number%10)

	fmt.Println("multiplication:", (number/10)*(number%10))

}
