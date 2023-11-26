package main

import (
	"fmt"
	"math"
)

func main(){
   var n,x int
   var min float64
   fmt.Println("Введите размер массива: ")
   fmt.Scan(&n)
   array:=make([]int,n)

   fmt.Println("Введите элементы массива: ")
   for i:=0;i<n;i++{
	fmt.Scan(&array[i])
   }

   fmt.Println("Введите целое число: ")
   fmt.Scan(&x)

   for i:=0;i<n;i++{
      min=math.Abs(float64(x-array[i]))
	  break
   }
   
   index :=0
   for i:=0; i < n;i++{
	if math.Abs(float64(x-array[i]))<min{
		min=math.Abs(float64(x-array[i]))
		index=i
	}

   }

   fmt.Println("Ответ: ",array[index])

}