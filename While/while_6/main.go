package main

import "fmt"

func main(){
	var n int
	var k float64
	fmt.Scan(&n)
	k=float64(n)
	if n%2!=0{
	for n>1{
	  n-=2
	  k*=float64(n)
	}
    }else{
	for n>2{
		n-=2
		k*=float64(n)
	  }
    }
	fmt.Println(k)
   
}