package main

import "fmt"

func readingNumber() {

	var sumEntered int
	fmt.Println("enter a target sum:")
	fmt.Scanln(&sumEntered)
	sum := 0
	count := 0

	for sum < sumEntered {
		var num int
		fmt.Print("enter a number:")
		fmt.Scanln(&num)

		sum += num
		count++
	}

	fmt.Printf("The sum of integers are equal or greater! Sum: %d, Number of integers entered: %d\n", sum, count)
}

//
//4.34 (Reading numbers until a specified sum) Write an application that asks for a number from
//the user and then keeps reading integer values from the user until the sum of those integers equals
//or becomes greater than the value that was input in the beginning.
