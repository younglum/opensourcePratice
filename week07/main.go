package main

import (
	"fmt"
	"strings"
)

func main() {
	// 오늘은 2024년 10월 11일 입니다.//
	var army string = "우리는 $가와 $민에 충성을 다하는 대한민$ 육군이다."
	armyFixed := strings.NewReplacer("$", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
	/*var now time.Time = time.Now()
	var mouth int = int(now.Month())
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.", now.Year(), mouth, int(now.Day()))
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.", now.Hour(), now.Minute(), now.Second())*/
}
