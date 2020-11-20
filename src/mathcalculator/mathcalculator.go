package mathcalculator

import (
	"fmt"
	"math"
)

func Absolute(n int) int{
	if n<0{
		return n * -1	// 입력받은 값이 음수일 경우 -1을 곱하여 절대값 반환
	} else{
		return n		// 입력받은 값이 양수일 경우 그대로 반환
	}
}

func Factorial(n int) int{
	if n == 0 {			// 입력받은 값이 0인 경우 1 반환
		return 1
	}
	return n * Factorial(n-1)	// 입력받은 값의 팩토리얼을 구하기 위해 n * (n-1)의 연산 후 반환 (재귀함수 사용)
}

func Exponentiation(n1 int, n2 int) int{
	return int(math.Pow(float64(n1), float64(n2)))	// 입력받은 두 값을 math.Pow() 함수를 이용하여 제곱근을 연산
	// 단, math.Pow() 함수는 정수형 숫자가 허용되지 않기에 형변환을 거쳐 연산 후 정수형으로 다시 캐스팅한다.
}

func Fibonacci(n int) int{
	if n <= 2 {
		return 1	// 입력받은 값이 2 이하일 경우 1을 반환한다.
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)	// 입력받은 값이 3 이상일 경우 (n-1) + (n-2)의 연산을 재귀함수를 이용하여 연산 후 반환한다.
	}
}
func Quadformula(n1 float64, n2 float64, n3 float64){
	var n4 = (n2 * n2) -(4 * n1 * n3)		// b2 - 4ac의 값을 판단하기 위한 변수이다.
	if n4 > 0 {								// b2 - 4ac의 값이 양수인 경우 (서로 다른 두 실근)
		fmt.Println("x1 값 : ",(-n2+(math.Sqrt(n4)))/2*n1)
		fmt.Println("x2 값 : ",(-n2-(math.Sqrt(n4)))/2*n1)
	} else if n4 == 0 {						// b2 - 4ac의 값이 0인 경우 (한 개의 중복되는 실근=중근)
		fmt.Println("x 값 : ", -n2/2*n1)
	} else if n4 < 0 {						// b2 - 4ac의 값이 음수인 경우 (해가 존재하지 않을 경우)
		fmt.Println("근이 존재하지 않습니다.")
	}
}

func Permutation(n1 int, n2 int) int{
	return Factorial(n1) / Factorial(n1-n2)		// nPr의 형태이므로 피보나치 함수를 이용하여 n! / (n-r)!의 연산 후 반환
}

func Percent(n1 float64, n2 float64) int{
	return int(n2 / n1 * 100)			// 입력받은 값이 float64형이기때문에 반환 시에는 정수형으로 다시 캐스팅한다.
}

func Prime(n int) int{
	if n <= 1 {return 0}		// 입력받은 값이 1보다 작거나 같은 경우 소수가 아님을 판단하는 0을 반환

	for i := 2; i < n; i++ {  	// 입력받은 값이 2에서 (입력받은 값-1)까지 나머지 연산을 수행하여 한번이라도
			if n % i == 0 { 	// 나머지 값이 0이 되는 경우가 있을 경우 소수가 아님을 판단하는 0을 반환
				return 0
			}
		}
	return 1					// 그 이외의 경우는 소수임을 판단하는 1을 반환한다.
}

func NCR(n1 int, n2 int) int{ 	// 조합계산은 nCr의 형태이므로 입력받은 두 정수를 Factorial 함수를 이용하여 n! / r! * (n-r)!의 연산을 수행한 후 반환한다.
	return Factorial(n1) / (Factorial(n2) * (Factorial(n1-n2)))
}

