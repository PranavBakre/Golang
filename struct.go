package main

import "fmt"

type vertex struct {
 a,b int
} 

func main(){
	var a vertex
	fmt.Println("Enter the value of a.a and a.b")
	fmt.Scanln(&a.a,&a.b)
	fmt.Println(a.a," ",a.b," ",a)
}
