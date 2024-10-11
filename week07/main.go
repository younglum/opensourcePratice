package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("점수를 입력하세요.")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                //줄바꿈을 제거, 파이썬의 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 10, 32) //10진 정수형(32비트)으로 변환
	if score >= 60 {
		status := "passing"
		fmt.Println(status)
		fmt.Printf("%d", score)
	} else {
		status := "failing"
		fmt.Println(status)
		fmt.Printf("%d", score)
	}

	/*var army string = "우리는 $가와 $민에 충성을 다하는 대한민$ 육군이다."
	armyFixed := strings.NewReplacer("$", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))*/
	/*var now time.Time = time.Now()
	var mouth int = int(now.Month())
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.", now.Year(), mouth, int(now.Day()))
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.", now.Hour(), now.Minute(), now.Second())*/
}
