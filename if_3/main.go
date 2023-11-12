package main

import "fmt"

func main(){
	var a int
    fmt.Scan(&a)
	if a > 0 {
		a+=1
		fmt.Println(a)
	}else if a < 0 {
		a-=2
		fmt.Println(a)
	}else{
		fmt.Println(10)
	}
	
}