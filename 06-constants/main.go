package main

import (
	"fmt"
	"math"
)

// const is read only. supports char, string, bool, numeric
// const can be declared everywhere as variable can
const str string = "constant"

func main() {
    fmt.Println(str)

	// numeric const has no type until it's given
    const num = 500000000

	// const can be declared with arbitrary precision
    const num2 = 3e20 / num
    fmt.Println(num2)

	// giving num2 int64 type
    fmt.Println(int64(num2))

	// giving num float64 as the "x" param in math.Sin() is float64. -> math.Sin(x float64){ ... }
    fmt.Println(math.Sin(num))
}
