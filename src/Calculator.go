package main

import (
	"fmt"
)


func Calcul(n1 int, n2 int, n3 int) int{
	return (n1+(n2*4)+n3)/6
}

func main(){
	var input int
	var input2 int
	var input3 int

	for {
		fmt.Println("PERT를 이용한 예상 시간 계산 작업을 수행합니다.")
		fmt.Print("낙관적 예상 시간을 입력하세요 : ")
		fmt.Scanln(&input)
		fmt.Print("현실적 예상 시간을 입력하세요 : ")
		fmt.Scanln(&input2)
		fmt.Print("비관적 예상 시간을 입력하세요 : ")
		fmt.Scanln(&input3)
		fmt.Println("활동이 완료되는 데까지 걸리는 예상시간은", Calcul(input, input2, input3), "입니다.")
		fmt.Println()
	}
}
