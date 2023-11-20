package main

import "fmt"

func main(){
	var a,b int
	var agree bool
    fmt.Scan(&a, &b)
	agree = a >=0 && b < -2
	fmt.Println(agree)
	
}