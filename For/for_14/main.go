package main

import (
	"fmt"
)

func main(){
   var n,sum int
   sum=0

   fmt.Println("Введите N:")
   fmt.Scan(&n)
   for i:=1; i<=(2*n-1);i+=2{
	 sum+=i
   }
   fmt.Println(sum)
}