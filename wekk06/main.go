package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13   //var i int = 13
	f := 12.8 //var f float64=12.8
	c1 := `2`
	c2 := `김` //44608,유니코드
	var t float64
	var a int
	var o bool
	mySchoolAccount := 5.0 //Camelcase 표기법
	fmt.Println(t, a, o, mySchoolAccount)
	fmt.Printf("%f %d %t\n", t, a, o) //zero value
	//fmt.Print("value i: %d,value f:%f", i, f)
	fmt.Printf("value i: %d,value f:%f\n", i, f)

	//fmt.Print("%d * %f = %f\n", i, f, i*f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2))
	fmt.Println(c1, c2)

	mySchoolAccount = 2.7
	c := 3
	fmt.Print("\n\n", mySchoolAccount > float64(c)) //비교연산 (true/false)

}
