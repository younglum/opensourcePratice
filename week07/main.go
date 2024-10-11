package main

import (
	"fmt"
	"time"
)

func main() {
	// 오늘은 2024년 10월 11일 입니다.//
	var now time.Time = time.Now()
	var mouth int = int(now.Month())
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.", now.Year(), mouth, int(now.Day()))
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.", now.Hour(), now.Minute(), now.Second())
}
