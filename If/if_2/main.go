package main

import "fmt"

func main(){
	var a int
    fmt.Scan(&a)
	if a > 0 {
		a+=1
		fmt.Println(a)
	}else{
		a-=2
		fmt.Println(a)
	}
	
}