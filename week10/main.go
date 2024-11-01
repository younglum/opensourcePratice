package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var counts int = 0

func isPrime(n int) bool {
	//var isPrime bool = true flag변수 , 사용시 가독성과 메모리 용량의 효율성을 챙길 수 있음
	if n < 0 {
		return false
	} else if n == 1 || n == 0 { //입력받은 값이 1보다 작거나 같을때 == 1 또는 음수일 때
		//isPrime = false
		fmt.Printf("1또는 0이 입력되었습니다.\n")
		return false
	} else if n == 2 {
		//isPrime = true
		return true
	} else if n%2 == 0 { //2를 제외한 짝수는 모두 소수가 아님,그래서 짝수일 때 무조건 소수가 아님을 출력
		//isPrime = false
		return false
	} else {
		j := 3
		//for j <= int(math.Sqrt(float64(n)))
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j) //for문의 반복횟수 출력
			j += 2
			counts++
		}
	}
	return true
}

func main() {
	//fmt.Printf("%f\n", math.Sqrt(16.8)) //16.8에 루트를 씌운 값을 출력
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your number: ")
	i, err := in.ReadString('\n') //입력을 문자열로 받음
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	//counts := 0
	/*j := 2
	for j < n {
		if n%j == 0 {
			//counts++
			isPrime = false
		}

		j++
	}*/
	if isPrime(n) {
		//if counts == 0 {
		fmt.Printf("%d is prime number.\n count : %d", n, counts)
	} else {
		fmt.Printf("%d is not prime number\n count : %d", n, counts)
		//fmt.Printf("%d is not prime number\n count : %d", n, counts)
	}
}
