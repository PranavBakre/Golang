package main
import (
	"facto/facto"
	"fmt")



func main(){
	var a,prod int64
	fmt.Printf("Enter a number\n")
	fmt.Scanf("%d",&a)
	fact a
	prod=a.Facto(a)
	fmt.Println(prod)		

}


