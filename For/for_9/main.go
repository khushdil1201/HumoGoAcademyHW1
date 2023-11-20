package main

import (
	"fmt"
	"math"
)

func main(){
   var a,b float64
   var k float64
   k=0
   fmt.Println("Введите A:")
   fmt.Scan(&a)
   fmt.Println("Введите B:")
   fmt.Scan(&b)
   for i:=a; i<=b;i++{
	k+=math.Pow(float64(i),2)
   }
   fmt.Println(k)
}