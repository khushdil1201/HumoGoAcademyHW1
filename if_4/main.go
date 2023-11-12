package main

import "fmt"

func main(){
	var a,b,c,temp int
	temp = 0
    fmt.Scan(&a, &b, &c)
    if a > 0{
		temp+=1
	}
	if b > 0{
		temp+=1
	}
	if c > 0{
		temp+=1
	}
	fmt.Println(temp)
	
	
}