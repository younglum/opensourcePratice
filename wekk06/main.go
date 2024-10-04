package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 13
	f := 12.8
	//fmt.Print("value i: %d,value f:%f", i, f)
	fmt.Printf("value i: %d,value f:%f\n", i, f)

	//fmt.Print("%d * %f = %f\n", i, f, i*f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i))
}
