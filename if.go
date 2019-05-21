package main

import "fmt"
import "runtime"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &num)
	if num > 10 {
		fmt.Printf("Hello")
	} else if num < 0 {
		fmt.Println("Bye")
	} else {
		fmt.Println("Bla Bla Bla")
	}
	
	fmt.Println("\n",runtime.GOARCH)
}
