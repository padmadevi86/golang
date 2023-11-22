package main

import "fmt"

func main() {

	var num int
	fmt.Println("Enter a number : ")
	fmt.Scanf("%d", &num)
	if n :=num /2; num%2 ==0{
		fmt.Printf("%d is Even ,%d times", num, n)
	}else {
		fmt.Printf("%d  is Odd, %d times", num, n)
	}
}



