// boiling 输出冰点的华氏和摄氏温度
package main

import "fmt"

// 冰点 华氏温度
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = fToC(f)
	fmt.Printf("boiling point = %g° F or %g° C \n", boilingF, c)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
