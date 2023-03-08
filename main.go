package myfirstpkg

import (
	"fmt"
)

func myFirstFunc() {
	fmt.Println("이건 사용 안됨 내부용")
}

func MyFirstFunc() {
	fmt.Println("내가 만든 첫번째 패키지 외부용")
}
