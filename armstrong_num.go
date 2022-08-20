package main

import "fmt"

func main() {
	Armstrong(371)
}

func Armstrong(num int) {
	//num := 371
	var sum int = 0
	var rem int = 0
	temp := num
	for num > 0 {
		rem = num % 10
		cube := rem * rem * rem

		sum = sum + cube
		num = num / 10

	}
	if sum == temp {
		fmt.Println("Number is Armstrong")

	} else {
		fmt.Println("Not an Armstrong ")

	}

}
