package main

import (
	"fmt"
)

func main(){
   var n int 
   var sum float64
   sum=0
   fmt.Println("Введите N:")
   fmt.Scan(&n)
   for i:=1; i<=n;i++{
	sum+=1/(float64(i))
   }
   fmt.Println(sum)
}