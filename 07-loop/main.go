package main

import "fmt"

func main() {
	// "for" is the only looping construct in Go

	// just like "for" in other languages, but without parentheses
	for i := 0; i < 3; i++ {
        fmt.Println(i)
    }

	// it can be act like "while" in other languages
    j := 1
    for j <= 3 {
        fmt.Println(j)
        j = j + 1
    }

	// we can use "range", it starts with 0
    for i := range 3 {
        fmt.Println("range", i)
    }

	// can run without init values. this will loop until you "break"
    for {
        fmt.Println("loop")
        break
    }

	// we can also continue to the next iteration of the loop
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
