package main

import "fmt"

func main(){
	var a,b,c int
	var agree bool
    fmt.Scan(&a, &b, &c)
	agree = a < b && b < c
	fmt.Println(agree)
	
}