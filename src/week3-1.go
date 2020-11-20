package main

import "fmt"

func fact(n int)int{
	result := 1
	for i := 2; i<=n; i++{
		result *= i
	}
	return result
}

func main(){
	fmt.Println(fact(6))


	//seconds := time.Now().Unix()
	//rand.Seed(seconds)
	//target := rand.Intn(100) + 1
	//fmt.Println(target)
}
