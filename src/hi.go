package main

import (
	"fmt"
	"reflect"
)

func main() { // Enrty Point
	var id, z int	// 변수 선언
	var x, y = 3.14, 2.71	// 타입 생략
	var human bool = true	// 선언과 동시에 초기화
	var a string	// 값을 할당하지 않고 선언한 변수는 해당 타입의 Zero 값으로 출력된다.
	var b bool		// 값을 할당하지 않고 선언한 변수는 해당 타입의 Zero 값으로 출력된다.

	id = 55
	//id = true		// 선언한 변수의 타입을 바꿀 수 없다.
	//x = 0
	//y = 0
	//x, y = 0, 10.8
	name := "trevor"	// C++의 Auto, Python처럼 변수 선언과 동시에 타입 지정 없이 초기화가 가능
	/*
	C++ 문법
	auto name = "trevor";	// string

	Python 문법
	name = "trevor"
	 */
	fmt.Println(a)	// 문자열의 Zero 값, 빈 문자열 출력
	fmt.Println(b)	// Bool 타입의 제로 값, false 출력
	fmt.Println(z)	// 숫자 타입의 제로 값, 0 출
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(human)
	fmt.Println(id)
	fmt.Println(x, y)
}
