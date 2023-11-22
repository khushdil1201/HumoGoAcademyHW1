package main

import "fmt"

func main(){
	var n int
	var k int
	fmt.Scan(&n)
	k=1
	for k*k<=n {
      k++
	}
	fmt.Println(k)
   
}