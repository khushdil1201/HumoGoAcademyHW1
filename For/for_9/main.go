package main

import (
	"fmt"
	"math"
)

func main(){
   var a,b int
   var k float64
   k=0
   fmt.Println("Введите A:")
   fmt.Scan(&a)
   fmt.Println("Введите B:")
   fmt.Scan(&b)
   for i:=a; i<=b;i++{
	k+=math.Pow10(i)
   }
   fmt.Println(k)
}