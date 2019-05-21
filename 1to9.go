package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the number\n")
	fmt.Scanf("%d", &a)
	switch a {
	//case 1:
	//fmt.Println("One")
	//break;

	case 0, 1: //test
		fmt.Println("Zero")
	//break;
	case 2:
		fmt.Println("Two")
	//break;
	case 3:
		fmt.Println("Three")
	//break;
	case 4:
		fmt.Println("Four")
		//		break;
	case 5:
		fmt.Println("Five")
		//		break;
	case 6:
		fmt.Println("Six")
		//		break;
	case 7:
		fmt.Println("Seven")
		//		break;
	case 8:
		fmt.Println("Eight")
		//		break;

	case 9:
		fmt.Println("Nine")
		//		break;

	default:
		fmt.Println("Error.")
	}

}
