package main

import (
	"fmt"
	"log"
	"../datafile"
)

func main(){
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(lines)
	}
	counts := make(map[string]int)	// Map
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
