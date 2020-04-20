package main

import "fmt"

func main() {
	//Pointers
	//Creating a pointer
	var firstname *string = new(string)
	*firstname = "Arthur"
	fmt.Print("Pointer: ")
	fmt.Println(firstname)
	//dereferencing
	fmt.Println("Value: " + *firstname)

	lastName := "Freeman"
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	var count int
	addPoint(&count)
	addValue(count)

	//swaping pointers
	a, b := 5, 10
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

//Creating a normal function
func addValue(count int) {
	count += 5
	fmt.Println("add5Value :", count)
}

//Creating a function with a pointer
func addPoint(count *int) {
	*count += 5
	fmt.Println("add5Value :", *count)
}

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
