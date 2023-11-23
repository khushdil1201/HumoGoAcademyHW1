package main

import "fmt"

func main(){
   var n int 
   fmt.Scan(&n)
   array:=make([]int, n)
   for i:=0;i<n;i++{
	  fmt.Scan(&array[i])
   }
   k:=0
   for i:=1;i<n-1;i++{
	  if array[i]>array[i-1] && array[i]>array[i+1]{
		k++
	  }
   }
   fmt.Println(k)

}