package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0; x < 8; x++ {
		f := math.Exp(float64(x))
		fmt.Printf("x = %d e^x = %8.5g e^x = %8.5[2]e e^x = %8.5[2]f\n", x, f)
	}
	fmt.Println(0.3 == 0.1+0.2)
}
