package main

import (
	"fmt"
	"time"
)

func main() {
	var date int
	var month time.Month
	var year int
	date = time.Now().Day()
	month = time.Now().Month()

	year = time.Now().Year()
	fmt.Printf("%vth %v%v", date, month, year)
}
