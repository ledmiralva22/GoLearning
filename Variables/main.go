package main

import (
	"fmt"
	"time"
)

var foo string = "bar"

func main() {
	//Long assigning
	var i int
	i = 42
	fmt.Println(i)

	//Short assigning
	var f float32 = 3.14
	//if the variable is not used, the compiler with throw an error
	fmt.Println(f)

	//Implicit initialization
	firstName := "Name"
	fmt.Println(firstName)

	//boolean variables
	b := true
	fmt.Println(b)
	b1 := 10
	fmt.Println(b1)
	b2 := "Last Name"
	fmt.Println(b2)

	//complex number
	c := complex(4, 3)
	fmt.Println(c)

	//multi assing variables
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//Excercises
	var baz string = "qux"
	fmt.Println(foo, baz)

	//Defining multivariables using var
	var (
		Debug       bool      = false
		LogLevel    string    = "info"
		startUpTime time.Time = time.Now()
	)
	fmt.Println(Debug, LogLevel, startUpTime)

	//Defining variable with no initial value
	var noValue int
	fmt.Println(noValue)

	//Defining multivariables using var without the type
	var (
		Debug2       = false
		LogLevel2    = "info"
		startUpTime2 = time.Now()
	)
	fmt.Println(Debug2, LogLevel2, startUpTime2)

	//Defining multivariables without the type
	Debug3, LogLevel3, StartupTime3 := false, "info", time.Now()
	fmt.Println(Debug3, LogLevel3, StartupTime3)

	//Assigning variables with functions
	Debug4, LogLevel4, StartupTime4 := getConfig()
	fmt.Println(Debug4, LogLevel4, StartupTime4)

	//Reusing a variable
	Debug4, LogLevel4, StartupTime4 = getConfig()
	fmt.Println(Debug4, LogLevel4, StartupTime4)
}

func getConfig() (bool, string, time.Time) {
	return true, "resume", time.Now()
}
