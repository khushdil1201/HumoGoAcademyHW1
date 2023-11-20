package main

import "fmt"

func main(){
   var a float64
   fmt.Println("Введите число:")
   fmt.Scan(&a)
   for i:=1.2; i<=2;i+=0.2{
	fmt.Println(a*i)
   }
}