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
	if n < 0 {
		return false
	} else if n == 1 || n == 0 {
		fmt.Printf("1또는 0이 입력되었습니다.\n")
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			j += 2
			counts++
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your start number: ")
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	bn := bufio.NewReader(os.Stdin)
	fmt.Print("input your end number: ")
	b, err := bn.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d is prime number.\n count : %d\n", j, counts)
		} else {
			fmt.Printf("%d is not prime number\n count : %d\n", j, counts)
		}
	}
}
