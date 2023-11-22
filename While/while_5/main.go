package main

import "fmt"

func main(){
	var n int
	k:=0
	fmt.Scan(&n)
	if n%2==0{
	   for n>=2 {
	       if n%2==0{
	       	k++
	       }else{
	       	k=0
	       }
	   	n/=2
	   }
   }
   if k==0{
	fmt.Println("Число не является степенью 2!")
   }else{
	fmt.Println(k)
   }
}