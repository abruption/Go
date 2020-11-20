package main

import (
	"fmt"
	"reflect"
)

func main(){
	//var id int = 55
	//var pid *int
	//pid = &id
	id := 55
	grade := 3.91
	pid := &id
	pgrade := &grade
	fmt.Println(grade, &grade, reflect.TypeOf(grade))	// &은 해당 변수의 주소 값, 실행할 때마다 달라짐
	fmt.Println(*pgrade, &pgrade, reflect.TypeOf(pgrade))
	fmt.Println(*pid, pid, &id)
	fmt.Println(*pid, pid, &pid)
}

/*
func main(){
	number := 3
	cube(number)
	fmt.Println(number)
}

func cube(n int){
	n = n * n * n
	fmt.Println(n)
}
 */