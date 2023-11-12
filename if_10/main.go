package main

import "fmt"

func main(){
	var a,b int
    fmt.Scan(&a, &b)
    if a != b{
		a= a+b
		b= a+b  
		fmt.Println("a = ",a," b = ",b)
	}else {
		a=0
		b=0
		fmt.Println("a = 0"," b = 0")
	}	
}	