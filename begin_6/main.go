package main

import "fmt"

func main() {
	var a,b,c,V,S float32

	fmt.Println("Enter a:")

	fmt.Scan(&a)

	fmt.Println("Enter b:")

	fmt.Scan(&b)

	fmt.Println("Enter c:")

	fmt.Scan(&c)

	V = a*b*c

	S = 2*((a*b)+(b*c)+(a*c))

	fmt.Println("V=",V)

	fmt.Println("S=",S)

	
}
