package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	//Initialize without an explicit array
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	//set a new item into to the slice
	slice2 = append(slice2, 4)
	fmt.Println(slice2)

	//create a new slice from the index 1 of the original slice
	slice3 := slice2[1:]
	//create a new slice until the index 2 of the original slice
	slice4 := slice2[:2]
	//create a new slice using the range especified in the brackets
	slice5 := slice2[1:2]
	//THIS IS A CHANGE

	fmt.Println(slice3, slice4, slice5)
}
