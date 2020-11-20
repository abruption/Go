package main

import (
	"fmt"
	"mathcalculator"
	"os"
)

func main(){
	var option int
	var input int
	var input2 int
	var input3 int

	for{
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 근의공식 6) 순열공식 7) 퍼센트 8) 소수판정 9) 조합공식 10) 종료 : ")
		fmt.Scanln(&option)		// 옵션 선택을 위한 숫자 입력

		if option == 1{		// 절대값 계산 옵션 선택
			fmt.Print("정수 입력(절대값 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(mathcalculator.Absolute(input)) // 입력 받은 값을 통해 절대값 함수 호출
		} else if option == 2{		// 팩토리얼 계산 옵션 선택
			fmt.Print("정수 입력(팩토리얼 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(mathcalculator.Factorial(input))	// 입력 받은 값을 통해 팩토리얼 함수 호출
		} else if option == 3{		// 피보나치 계산 옵션 선택
			fmt.Print("정수 입력(피보나치 계산) : ")
			fmt.Scanln(&input)
			fmt.Println(mathcalculator.Fibonacci(input))	// 입력 받은 값을 통해 피보나치 함수 호출
		} else if option == 4{		// 거듭제곱(제곱근) 계산 옵션 선택
			fmt.Print("Base 입력 : ")
			fmt.Scanln(&input)
			fmt.Print("Exponent 입력 : ")
			fmt.Scanln(&input2)
			fmt.Println(mathcalculator.Exponentiation(input, input2))	// 입력 받은 값을 통해 거듭제곱(제곱근) 함수 호출
		} else if option == 5{		// 근의공식 계산 옵션 선택
			fmt.Println("ax2 + bx + c = 0의 해를 구하는 메뉴입니다.")
			fmt.Print("a 값 입력 : ")
			fmt.Scanln(&input)
			fmt.Print("b 값 입력 : ")
			fmt.Scanln(&input2)
			fmt.Print("c 값 입력 : ")
			fmt.Scanln(&input3)
			mathcalculator.Quadformula(float64(input), float64(input2), float64(input3))	// 입력 받은 값을 float64로 캐스팅하여 근의공식 함수 호출
		} else if option == 6{		// 순열계산 계산 옵션 선택
			fmt.Print("정수 값 입력(순열 계산) : ")
			fmt.Scanln(&input)
			fmt.Print("정수 값 입력 : ")
			fmt.Scanln(&input2)
			fmt.Println(mathcalculator.Permutation(input, input2))	// 입력 받은 값을 통해 순열계산 함수 호출
		} else if option == 7{		// 퍼센트 계산 옵션 선택
			fmt.Print("분모 값 입력(퍼센트 계산) : ")
			fmt.Scanln(&input)
			fmt.Print("분자 값 입력 : ")
			fmt.Scanln(&input2)
			fmt.Println(mathcalculator.Percent(float64(input), float64(input2)),"%") 	// 입력 받은 값을 나누기 연산을 위해 float64로 캐스팅하여 퍼센트 함수 호출
		} else if option == 8{ 		// 소수판정 옵션 선택
			fmt.Print("정수 입력(소수판정 확인) : ")
			fmt.Scanln(&input)
			var result = mathcalculator.Prime(input)	// 입력 받은 값을 통해 소수판정 함수 호출 후 값을 반환받아 result 변수에 저장한다.
			if result == 0{		// 반환받은 값이 0인 경우 입력 받은 값이 소수가 아님을 출력한다.
				fmt.Println(input, "은(는) 소수가 아닙니다.")
			} else{				// 반환받은 값이 0인 아닌 경우 입력 받은 값이 소수임을 출력한다.
				fmt.Println(input,"은(는) 소수입니다.")
			}
		} else if option == 9{		// 조합계산 옵션 선택
			fmt.Print("원소의 개수 입력(조합 계산) : ")
			fmt.Scanln(&input)
			fmt.Print("경우의 수 입력 : ")
			fmt.Scanln(&input2)
			fmt.Println(mathcalculator.NCR(input, input2))		// 입력 받은 값을 통해 조합계산 함수 호출
		} else if option == 10{
			os.Exit(3)	// 프로그램 종료
		} else {
			fmt.Println("잘못 된 입력 값입니다. 1~10 사이의 수를 입력하세요.")	// 오류 방지
		}
	}
}
