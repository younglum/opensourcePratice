package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	fmt.Println(scores)
	fmt.Printf("%#v\n", scores) //#v하면 배열의 리터럴을 그대로 보여준다

	// 	var dates [3]time.Time
	// 	dates[1] = time.Unix(1447920000, 0)
	// 	fmt.Print(dates[1])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(14479200000, 0)}

	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Printf("%#v\n", dates)

}
