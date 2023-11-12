package main

import "fmt"

func main(){
	var a int
	var agree bool
    fmt.Scan(&a)
	agree = a % 2== 0
	fmt.Println(agree)
	
}