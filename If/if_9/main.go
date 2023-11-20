package main

import "fmt"

func main(){
	var a,b float64
    fmt.Scan(&a, &b)
    if a > b{
		b,a = a,b
		fmt.Println("a=",a," b=",b)
	}else {
		fmt.Println("a=",a," b=",b)
	}	
}	