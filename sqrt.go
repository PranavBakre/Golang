package main

import (
	"fmt"
)

func Sqrt(x float32) float32 {
	z := x
	fmt.Println(z)
	
	for z*z != x {
		z -= (z*z - x)/(2*z) 
		
	}

	return z
}

func main() {
	var x float32
	
	fmt.Println("Enter the number")
	fmt.Scanf("%g",&x)
	fmt.Println(Sqrt(x))
}

