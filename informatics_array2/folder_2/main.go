package main

import "fmt"

func main(){
   var n,x int
   fmt.Println("Введите размер массива: ")
   fmt.Scan(&n)
   array:=make([]int,n)

   fmt.Println("Введите элементы массива: ")
   for i:=0;i<n;i++{
	fmt.Scan(&array[i])
   }

   fmt.Println("Введите целое число: ")
   fmt.Scan(&x)

   count:=0
   for i:=0; i < n;i++{
      if array[i]==x{
		count++
	  }
   }
   if count>0{
	   fmt.Println("YES")
   }else{
	   fmt.Println("NO")
   }


}