package main

import (
	"bufio"
	"fmt"
	"os"
)

//第一个版本的dup输出标准输入流中的出现多次的行，在行内容前是出现次数的计数。这个程序将引入if 表达式，map内置数据结构和bufio的package。

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
		if input.Err() != nil {
			panic("input is error")
		}
	}

	for line, n := range count {
		if n <= 1 {
			continue
		}
		fmt.Printf("%d\t%s\n", n, line)
	}
}
