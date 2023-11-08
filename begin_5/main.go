package main

import "fmt"

func power(a float32, n int) float32 {
	if n == 1 {
		return a
	} else {
		return a * power(a, n-1)
	}
}

func main() {
	var a, V, S float32

	fmt.Println("Enter a:")

	fmt.Scan(&a)

	V = power(a, 3)

	S = 6 * power(a, 2)

	fmt.Println("V=", V)

	fmt.Println("S=", S)

}
