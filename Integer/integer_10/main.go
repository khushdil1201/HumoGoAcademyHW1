package main

import "fmt"

func main() {
	var number,sum,multy int

	fmt.Println("Enter number:")

	fmt.Scan(&number)

	sum = (number/100)+(number%10)+((number/10)%10)

	multy = (number/100)*(number%10)*((number/10)%10)

	fmt.Println("sum: ", sum)

	fmt.Println("multiplication: ", multy)
}
