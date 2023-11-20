package main

import (
	"fmt"
)

func main(){
   var n int
   var p,a float64
   p=1
   a=1.1
   fmt.Println("Введите N:")
   fmt.Scan(&n)
   for i:=0; i<=n;i++{
	p*=a
	a+=0.1
   }
   fmt.Println(p)
}