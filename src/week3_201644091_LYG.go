package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){

	fmt.Print("첫 번째 정수 입력 : ")
	reader := bufio.NewReader(os.Stdin)	// 공백 문자까지 포함하여 입력 받음
	input1, err := reader.ReadString('\n')
	fmt.Print("두 번째 정수 입력 : ")
	input2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)	// 에러처리
	}

	var in1 = strings.TrimSpace(input1)	// 문자열 앞뒤에 오는 공백 문자 제거 (줄바꿈 문 제거)
	var in2 = strings.TrimSpace(input2)	// 문자열 앞뒤에 오는 공백 문자 제거 (줄바꿈 문 제거)
	one, err := strconv.Atoi(in1)	// 문자열을 integer 타입으로 변환
	two, err := strconv.Atoi(in2)	// 문자열을 integer 타입으로 변환
	if err != nil {
		log.Fatal(err)	// 에러처리
	}

	for i := 1; i<=50; i++{
		if i % one != 0 && i % two != 0{
			fmt.Println(i)	// 입력받은 두 숫자의 나머지 값이 0이 아닌 경우 출력
		} else if i % one == 0 && i % two == 0 {
			fmt.Println(one, two, "의 공배수", i)	// 입력받은 두 숫자의 나머지 값이 0인 경우 출력 (공배수)
		} else if i % one == 0 {
			fmt.Println(one, "의 배수")	// 입력받은 첫 번째 숫자의 나머지 값이 0인 경우 출력
		} else if i % two == 0 {
			fmt.Println(two, "의 배수")	// 입력받은 두 번째 숫자의 나머지 값이 0인 경우 출력
		}
	}
}