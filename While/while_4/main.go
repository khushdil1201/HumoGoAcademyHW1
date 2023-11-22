package main

import "fmt"

func main(){
	var n int
	var b bool
	fmt.Scan(&n)
	if n%3==0{
	  for n>3{
         n/=3 
	  }
    }
	if n==3{
		b=true
	}else{
		b=false
	}
	fmt.Println(b)
}