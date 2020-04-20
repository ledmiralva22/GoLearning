package main

import "fmt"

func main() {
	//Main Course
	var total float64 = 2 * 13
	fmt.Println("Sub: ", total)

	//Drinks
	total = total + (4 * 2.25)
	fmt.Println("Sub: ", total)

	//Discount
	total = total - 5
	fmt.Println("Sub: ", total)

	//Tip
	tip := total * 0.1
	fmt.Println("Tip: ", tip)

	//Total
	total = total + tip
	fmt.Println("Total: ", total)

	//Split
	split := total / 2
	fmt.Println("Split: ", split)

	//Reward
	visitCount := 24
	visitCount++
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("You earned a reward")
	}

	//Join strings
	firstName := "Luis"
	LastName := "Miranda"
	fmt.Println(firstName + " " + LastName)

	var count int
	fmt.Printf("Count : %#v \n", count)
}
