package main

import (
	"fmt"
	"os"
	"reflect"
)

// func Test(strs string) { //하나밖에 못받음
func Test(strs ...string) { //...은 가변매게변수를 처리할 때 사용,출력시 슬라이스 형태로 출력됨, 마지막 변수에만 사용가능
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	//fmt.Println(os.Args[1:], len(os.Args))
	Slice := os.Args[1:] //실행하는 파일이름을 출력하지 않게 자름
	fmt.Println(Slice[1])
	for _, slice := range Slice {
		fmt.Println(slice)
	}
	Slice = append(Slice, "Forever", "happy") //append({슬라이스},{추가할 숫자 또는 문자}.......)
	fmt.Printf("%s  %d", Slice, len(Slice))
	Test("abc")
	Test("a", "b")

	/*
		var emptySliace []bool
		//emptySliace = make([]bool, 0)
		fmt.Printf("%#v %d\n", emptySliace, len(emptySliace))

		if len(emptySliace) == 0 { //emptySlice의 크기가 0이면 실행. 선언하고 아무것도 안 넣었기 때문에 0
			emptySliace = append(emptySliace, true)
		}
		fmt.Printf("%#v  %d\n", emptySliace, len(emptySliace))


			var arrays [5]float64 = [5]float64{3.5, 4.1, 4.5, 5.0, 4.2}
			arrays_slice := arrays[1:4] //슬라이스 생성
			//arrays_slice[1] = 2.71
			arrays[2] = 2.71 //어디서 바꾸든 둘 다 출력이 동일해짐=슬라이싱은 array를 포인터로 가르키는것과 동일
			arrays_slice = append(arrays_slice, 4.6)
			fmt.Println(arrays, reflect.TypeOf(arrays))
			fmt.Println(arrays_slice, reflect.TypeOf(arrays_slice))
	*/
}
