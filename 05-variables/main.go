package main

import "fmt"

func main()  {
	// variable declaration
	// we can use "var"

	var emptyString string // default to ""
	var defaultInt int // default to 0
	var defaultFloat float32 // default to 0
	var defaultBool bool // default to false

	var myString string
	myString = "Hi!"

	// declare and initialize. data type is automatically detected, like a magic!
	var myString2 = "Hi, mom!"

	// or use := as a shorthand to declare and initialize value
	myString3 := "Hi, dad!"

	// we also can declare and init multiple vars at time
	num1, num2, num3 := 90, 6, 99

	fmt.Println(emptyString)
	fmt.Println(defaultInt)
	fmt.Println(defaultFloat)
	fmt.Println(defaultBool)
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)
	fmt.Println(num1, " + ", num2, " = ", num3) // context: jokes aside, it's a controversial AFC Asian Qualifier match; Bahrain vs Indonesia, 20241011.
}
