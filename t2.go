package main

import "fmt"
type node struct {
	left,right *node
	data int
}

func main(){
	var choice int
	var choice2 string
	root := &node{}
	fmt.Println("Enter the root data")
	fmt.Scanln(&root.data)
	fmt.Println("Enter the choice")
	fmt.Scanln(&choice)
	for choice ==1 {
		temp:=root
		curr:=&node{}
		fmt.Println("Enter the node data")
		fmt.Scanln(&curr.data)
		for {
			fmt.Println("L/R?")
			fmt.Scanln(&choice2)
			fmt.Println(choice2)
			if choice2=="l" {
				if temp.left==nil {
					temp.left=curr
					break
				}else {
					temp=temp.left
				}
			}else if choice2=="r" {
				if temp.right==nil {
					temp.right=curr
					break
				}else {
					temp=temp.right
				}
			}else{
				fmt.Println("Error")
			}
			
			
		}
		fmt.Println("Continue?")
		fmt.Scanln(&choice)
	}
	temp:=root
	for temp!=nil {
	fmt.Println(temp.data)
	temp=temp.left
	}
}
