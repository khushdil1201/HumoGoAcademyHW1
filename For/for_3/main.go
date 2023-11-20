package main

import "fmt"

func main(){
   var a,b,k int
   fmt.Println("A=")
   fmt.Scan(&a)
   fmt.Println("B=")
   fmt.Scan(&b)
   for i:=b-1; i>a;i--{
	fmt.Println(i)
	k=b-a-1
   }
   fmt.Println(k)
}