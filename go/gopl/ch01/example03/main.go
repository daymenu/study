package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 输出的是
	fmt.Println(os.Args[1:])

	fmt.Printf("%v", os.Args[1:])

	// 输出的是一个字符串
	fmt.Println(strings.Join(os.Args[1:], " "))
}
