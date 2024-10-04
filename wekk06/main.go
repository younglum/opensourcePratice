package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13   //var i int = 13
	f := 12.8 //var f float64=12.8
	c1 := `2`
	c2 := `김`
	//fmt.Print("value i: %d,value f:%f", i, f)
	fmt.Printf("value i: %d,value f:%f\n", i, f)

	//fmt.Print("%d * %f = %f\n", i, f, i*f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2))
	fmt.Println(c1, c2)
}
