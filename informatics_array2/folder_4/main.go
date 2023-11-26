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

   j:=0
   for i:=0; i < n;i++{
	   if array[i]==x{
		   j++
		}
	}

	a:=make([]int, j)
	k:=0
    for i:=0; i < n;i++{
		if array[i]==x{
			a[k]=i+1
			k++
		 }
	 }


   fmt.Println("Ответ: ")
   for i:=0; i < j;i++{
	  fmt.Println(a[i])
 }
   

}