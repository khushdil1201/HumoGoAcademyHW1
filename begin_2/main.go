package main

import "fmt"

func main() {
	var number float32

	fmt.Println("Введите число:")

	fmt.Scan(&number)

	fmt.Println("Answer:", number*number)

}
