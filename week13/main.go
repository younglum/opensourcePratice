package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arrays [5]float64 = [5]float64{3.5, 4.1, 4.5, 5.0, 4.2}
	arrays_slice := arrays[1:4] //슬라이스 생성
	//arrays_slice[1] = 2.71
	arrays[2] = 2.71 //어디서 바꾸든 둘 다 출력이 동일해짐=슬라이싱은 array를 포인터로 가르키는것과 동일
	arrays_slice = append(arrays_slice, 4.6)
	fmt.Println(arrays, reflect.TypeOf(arrays))
	fmt.Println(arrays_slice, reflect.TypeOf(arrays_slice))
}
