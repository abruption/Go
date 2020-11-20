package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {	// Enrty Point
	fmt.Println('C')	// 유니코드 출력
	fmt.Println('\n')
	fmt.Println(true)
	fmt.Println(5 != 5)
	fmt.Println("Hello, \"GO!\"")
	fmt.Println(math.Ceil(3.14))			// 소수점 올림으로 4가 출력됨
	fmt.Println(strings.Title("hello, go ~"))		// 각 문자의 첫번째를 대문자로 변환

}