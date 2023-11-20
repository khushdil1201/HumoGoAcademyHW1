package main

import (
	"fmt"
	"math"
)

func main(){
   var n int
   var sum float64
   sum=0
   fmt.Println("Введите N:")
   fmt.Scan(&n)
   for i:=0; i<=n;i++{
	sum+=math.Pow((float64(n)+float64(i)),2)
   }
   fmt.Println(sum)
}