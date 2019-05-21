package main

import "fmt"

func main() {

	var s1, s2, length int = -1, 1, 0
	var i, si int
	fmt.Scanln(&length)
	for i = 0; i < length; i++ {
		si = s1 + s2
		s1 = s2
		s2 = si
		fmt.Println(si)
	}

}
