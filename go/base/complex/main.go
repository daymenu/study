package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x * y)
	fmt.Printf("%g %[1]T\n", real(x*y))
	fmt.Printf("%g %[1]T\n", imag(x*y))

	var z complex64 = complex(5, 10)
	fmt.Println(z)
	fmt.Printf("%g %[1]T\n", real(z))
	fmt.Printf("%g %[1]T\n", imag(z))
}
