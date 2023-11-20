package main

import "fmt"

func main() {
	var L,R,S float32
    
	const p=3.14

	fmt.Println("Enter R:")

	fmt.Scan(&R)

	L = 2*p*R

	S = p*(R*R)

	fmt.Println("L=",L)

	fmt.Println("S=",S)

}
