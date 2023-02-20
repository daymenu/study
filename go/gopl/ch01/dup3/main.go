package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 我们也可以一次性的把整个输入内容全部读到内存中，然后再把其分割为多行，然后再去处理这些行内的数据。

func main() {
	counts := make(map[string]int)
	argc := len(os.Args[1:])

	if argc < 1 {
		fmt.Println("please input filename")
		return
	}

	for _, filePath := range os.Args[1:] {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("open filen %s is file: %v", filePath, err)
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
