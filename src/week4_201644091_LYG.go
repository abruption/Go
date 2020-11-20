package main

import (
	"fmt"
	"math"
	"os"
)

func absolute(n int) int{
	if n<0{
		return n * -1	// 입력받은 값이 음수일 경우 -1을 곱하여 절대값 반환
	} else{
		return n		// 입력받은 값이 양수일 경우 그대로 반환
	}
}

func factorial(n int) int{
	if n == 0 {			// 입력받은 값이 0인 경우 1 반환
		return 1
	}
	return n * factorial(n-1)	// 입력받은 값의 팩토리얼을 구하기 위해 n * (n-1)의 연산 후 반환 (재귀함수 사용)
}

func exponentiation(n1 int, n2 int) int{
	return int(math.Pow(float64(n1), float64(n2)))	// 입력받은 두 값을 math.Pow() 함수를 이용하여 제곱근을 연산
	// 단, math.Pow() 함수는 정수형 숫자가 허용되지 않기에 형변환을 거쳐 연산 후 정수형으로 다시 캐스팅한다.
}

func fibonacci(n int) int{
	if n <= 2 {
		return 1	// 입력받은 값이 2 이하일 경우 1을 반환한다.
	} else {
		return fibonacci(n-1) + fibonacci(n-2)	// 입력받은 값이 3 이상일 경우 (n-1) + (n-2)의 연산을 재귀함수를 이용하여 연산 후 반환한다.
	}
}

func main(){
	var option int
	var input int
	var input2 int

	for{
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 종료 : ")
		fmt.Scanln(&option)		// 옵션 선택을 위한 숫자 입력

		if option == 1{		// 절대값 계산 옵션 선택
			fmt.Print("정수 입력(절대값 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(absolute(input))	// 입력 받은 값을 통해 절대값 함수 호출
		}
		if option == 2{		// 팩토리얼 계산 옵션 선택
			fmt.Print("정수 입력(팩토리얼 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(factorial(input))	// 입력 받은 값을 통해 팩토리얼 함수 호출
		}
		if option == 3{		// 피보나치 계산 옵션 선택
			fmt.Print("정수 입력(피보나치 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(fibonacci(input))	// 입력 받은 값을 통해 피보나치 함수 호출
		}
		if option == 4{		// 거듭제곱(제곱근) 계산 옵션 선택
			fmt.Print("Base 입력 : ")
			fmt.Scanln(&input)
			fmt.Print("Exponent 입력 : ")
			fmt.Scanln(&input2)
			fmt.Println(exponentiation(input, input2))	// 입력 받은 값을 통해 거듭제곱(제곱근) 함수 호출
		}
		if option == 5 {
			os.Exit(3)	// 프로그램 종료
		}
	}
}
