package main

import "fmt"

func main() {
	var size int

	fmt.Println("Enter size:")

	fmt.Scan(&size)

	fmt.Println("Result:", size/1024)

}
