package main

import (
	"fmt"
	"time"
)

func main(){
	var dates [3]time.Time
	dates[0] = time.Unix(12357894000, 0)
	//dates[1] = time.Unix(62342313000, 1)
	//dates[2] = time.Unix(16234234000, 0)
	fmt.Println(dates[1])

	var sum int = 0
	primes := [5]int{2, 3, 5, 7, 11}
	for _, v := range primes{
		sum += v
	}
	//fmt.Println(float64(sum) / float64(len(primes)))
	fmt.Printf("%0.2f\n", float64(sum) / float64(len(primes)))

	//for k :=0; k<= len(primes); k++ {
	//	fmt.Println(primes[k])
	//}
	//fmt.Println(primes[1], primes[4])


	var notes [7]string
	notes[0] = "도"
	notes[1] = "레"
	notes[2] = "미"
	fmt.Print(notes[0], notes[2])
}