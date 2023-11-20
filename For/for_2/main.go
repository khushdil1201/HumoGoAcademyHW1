package main

import "fmt"

func main(){
   var a,b,k int
   fmt.Println("A=")
   fmt.Scan(&a)
   fmt.Println("B=")
   fmt.Scan(&b)
   for i:=a; i<=b;i++{
	fmt.Println(i)
	k=i-a+1
   }
   fmt.Println(k)
}