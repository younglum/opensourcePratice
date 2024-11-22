package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args[1:], len(os.Args))
	Slice := os.Args[1:] //실행하는 파일이름을 출력하지 않게 자름
	fmt.Println(Slice[1])
	for _, slice := range Slice {
		fmt.Println(slice)
	}
	Slice = append(Slice, "Forever", "happy")
	fmt.Printf("%s  %d", Slice, len(Slice))

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
