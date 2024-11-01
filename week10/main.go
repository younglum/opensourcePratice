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
	counts := 0
	j := 2
	for j < n {
		if n%j == 0 {
			counts++
		}

		j++
	}
	if counts == 0 {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number\n count : %d", n, counts)
	}
}
