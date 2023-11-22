package main

import "fmt"

func main(){
	var a,b float64
	fmt.Scan(&a,&b)
	for a-b>=0 {
	   	a-=b
	}
	fmt.Println(a)
}