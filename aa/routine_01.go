package main 

import (
	"fmt"
	"time"	// Java Thread sleep(5000)
)

func x(){
	for i:=0; i<10; i++ {
		fmt.Print("x")
	}
}

func y(){
	for i:=0; i<10; i++ {
		fmt.Print("y")
	}
}

func z(){
	fmt.Println("z")
}

func k(){
	fmt.Println("k")
}

func main(){
	defer z()	// 메인함수가 끝날때까지 지연시킴
	defer k()	// Stack의 LIFO 방식으로 동작 
	go x()
	go y()
	time.Sleep(time.Second * 2)
	fmt.Println()
	fmt.Println("end Main()")
}
