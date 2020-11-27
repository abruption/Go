package main

import "fmt"

func main(){
	//var ranks map[string]int
	//ranks = make(map[string]int)
	//ranks := make(map[string]int)
	ranks := map[string]int{"gold":1, "silver":2, "bronze":3}
	//ranks["gold"] = 1
	//ranks["silver"] = 2
	//ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])	
}
