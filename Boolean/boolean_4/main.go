package main

import "fmt"

func main(){
	var a,b int
	var agree bool
    fmt.Scan(&a, &b)
	agree = a > 2 && b <= 3
	fmt.Println(agree)
	
}