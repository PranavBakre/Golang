package main

import "fmt"

func main(){
var i int
var p *int
p=&i
fmt.Println("Enter the value of i")
fmt.Scanln(p)
fmt.Println(i," ",*p)
}
