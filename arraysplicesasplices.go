package main

import "fmt"

func main() {
	var array []int=make([]int,5)
	
	for i:=0;i<5;i++ {
		fmt.Scanln(&array[i])
	}
	
	var disp=fun(array)
	
	fmt.Println(disp)
}
func fun (arr []int) int {
	if arr[0] > 1 {
		return 1
	}else {
		return 0
	}
}
