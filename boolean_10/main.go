package main

import "fmt"

func main(){
	var a,b int
	var agree bool
    fmt.Scan(&a, &b)
	agree = (a % 2 != 0 && b % 2 == 0) || (a % 2 == 0 && b % 2 != 0)
	fmt.Println(agree)
	
}