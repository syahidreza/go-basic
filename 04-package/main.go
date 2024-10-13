package main // Go program is made up of packages. "main" is default start-up package in Go

import (
	"fmt"  // this is a package
	"math" // this is also a package
)

func main()  {
	// Pi is a const inside "math" package.
	// It begins with capital letter, means it is exported and can be accessed from outside of "math" package.

	// this won't print pi value
	// fmt.Println(math.pi)

	// this will print pi value
	fmt.Println(math.Pi)
}
