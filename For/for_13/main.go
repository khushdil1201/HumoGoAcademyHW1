package main

import (
	"fmt"
	"math"
)

func main(){
   var n int
   var p,a float64
   p=0
   a=1.1
   fmt.Println("Введите N:")
   fmt.Scan(&n)
   for i:=0; i<=n;i++{
	p+=a*(math.Pow(1,float64(i)))
	a+=(0.1)
   }
   fmt.Println(p)
}