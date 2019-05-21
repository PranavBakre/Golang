package main
import "fmt"

func main() {
	var principal,rate,time,interest int;
	fmt.Printf("Enter the principal amount,Rate and Time");
	fmt.Scanln(&principal,&rate,&time)
	interest=principal*rate*time/100;
	fmt.Println(interest);
	
}
