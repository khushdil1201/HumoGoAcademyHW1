package main

import "fmt"

func main(){
	var n int
	var k int
	fmt.Scan(&n)
	k=1
	for n>0 {
	 if (k*k<=n){
      k++
	 }else{
		break
	 }
	}
	fmt.Println(k-1)
   
}