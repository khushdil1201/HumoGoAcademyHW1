package main

import "fmt"

func main(){
	var a,b float64
	k:=0
	fmt.Scan(&a,&b)
	for a-b>=0 {
	   	a-=b
		k++
	}
	fmt.Println(k)
}