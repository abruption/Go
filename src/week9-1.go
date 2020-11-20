package main

import "fmt"

func main(){
	primeArray := [5]int{2,3,5,7,11}
	//primeSlice := []int{2,3,5,7,11}
	var primeSlice []int
	primeSlice = make([]int, 5)
	primeSlice[3] = 7

	//fmt.Println(primeArray)
	//fmt.Println(primeSlice)
	fmt.Println(primeArray[3])
	fmt.Println(primeSlice[3])
}