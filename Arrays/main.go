package main

import "fmt"

func main() {
	//Arrays
	//Long form
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[0])

	//Implicit initialization
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}
