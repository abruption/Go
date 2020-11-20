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
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}	
	if matched == false {
		names = append(names, line)	// 슬라이스 추가 append
		counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
