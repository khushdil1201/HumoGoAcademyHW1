package main

import "fmt"

func main(){
   var n int 
   fmt.Scan(&n)
   array:=make([]int, n)
   for i:=0;i<n;i++{
	  fmt.Scan(&array[i])
   }
   for i:=0;i<n/2;i++{
	  array[i],array[n-1-i]=array[n-1-i],array[i]
   }
   fmt.Println(array)

}