package main

import (
	"fmt"
)

func main(){
	var n, temp float64
	var k int
	temp=3
	fmt.Scan(&n)
	k=1
    for temp<=n{
	  temp*=3
     k++
	}
	fmt.Println(k)
   
}