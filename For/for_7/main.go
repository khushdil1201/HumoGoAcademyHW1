package main

import "fmt"

func main(){
   var a,b,k int
   k=0
   fmt.Println("Введите A:")
   fmt.Scan(&a)
   fmt.Println("Введите B:")
   fmt.Scan(&b)
   for i:=a; i<=b;i++{
	k+=i
   }
   fmt.Println(k)
}