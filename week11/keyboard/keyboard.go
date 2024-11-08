package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetIntager() (int, error) {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your start number: ")
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return n1, nil
}

func GetFloat() (float64, error) {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your start number: ")
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n1, nil
}
