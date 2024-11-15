package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
	"github.com/headfirstgo/keyboard"
)

func main() {
	var GP [3]float64
	var h error
	for i := 0; i < len(GP); i++ {
		fmt.Print("input Float number please : ")
		GP[i], h = keyboard.GetFloat()
		if h != nil {
			log.Fatal(h)
		}
	}
	for _, gpa := range GP {
		fmt.Println(gpa)
	}

	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	/*
		var scores [3]int
		scores[1] = 90
		//fmt.Println(scores)
		//fmt.Printf("%#v\n", scores) //#v하면 배열의 리터럴을 그대로 보여준다

		for i := 0; i < len(scores); i++ {
			fmt.Printf("%d\n", scores[i])
		}

		// 	var dates [3]time.Time
		// 	dates[1] = time.Unix(1447920000, 0)
		// 	fmt.Print(dates[1])

		dates := [3]time.Time{
			time.Unix(0, 0),
			time.Unix(1, 0),
			time.Unix(14479200000, 0)}

		fmt.Print("")
		// fmt.Println(dates[0], dates[1], dates[2])
		// fmt.Printf("%#v\n", dates)
		// for i, date := range dates {
		// 	fmt.Println(i, date)
		// }
		for _, date := range dates {
			fmt.Println(date)
		}*/

}
