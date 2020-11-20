package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 이름은 문자로 시작, 이후에는 문자/숫자 모두 사용 가
// 변수, 함수, 타입이 이름이 대문자로 시작하면 외부로 노출되어 외부 패키지에서 접근 가능함.
func main()  {
	a := 3
	a, b := 2, 7
	fmt.Println(a, b)


	//var append string = "shadow"		// 내장함수 이름으로 변수를 선언 시 변수로 동작하게 됨.
	var test = append([]string{}, "테스트")	// 내장함수로 동작.
	//fmt.Println(append)
	fmt.Println(test)


	fmt.Print("반지름 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)	// 줄바꿈 문 제거
	radius, err := strconv.ParseFloat(in, 64)	// 실수형으로 변환
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("원의 넓이 : " , radius * radius * 3.141592)



	//wrongTexts := "Gx Gx Gx!"
	//replacer := strings.NewReplacer("x", "o")
	//correctTexts := replacer.Replace(wrongTexts)
	//fmt.Println(correctTexts)

	//var now time.Time = time.Now()
	//var year int = now.Year()
	//month := now.Month()
	//fmt.Println(year)
	//fmt.Println(int(month))	// 캐스팅하지 않으면 문자로 출력 (ex. September)

	//var length float64 = 5.7
	//var width int = 5
	//length = width

	//fmt.Println("면적은 ", length*width)
	//fmt.Println(int(length))
	//fmt.Println("크기비교 ", int(length) >= width)
}
