package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var mouth int = int(now.Month())
	fmt.Printf("%d", mouth)
}
