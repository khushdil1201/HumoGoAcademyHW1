package main

import (
	"fmt"
)

func main(){
   var n int
   var a,pro float64
   pro=1
   fmt.Println("Введите N:")
   fmt.Scan(&n)
   fmt.Println("Введите A:")
   fmt.Scan(&a)
   for i:=1; i<=n;i++{
	 pro*=a
   }
   fmt.Println(pro)
}