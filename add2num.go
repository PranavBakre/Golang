package main

import "fmt"

func add(a int, b int) (c int) {

	c = a + b
	return
}

func main() {
	var a, b = 5, 6

	fmt.Println(add(a, b))
}
