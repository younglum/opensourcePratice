package main

import "fmt"

func main() {
	ages := make(map[string]int)
	var name string
	var age int

	for {
		fmt.Print("이름을 입력하세요 : ")
		fmt.Scanln(&name)
		if name == "q" || name == "Q" {
			break
		}
		fmt.Print("나이를 입력하세요 : ")
		fmt.Scanln(&age)

		ages[name] = age
	}
	for name, age := range ages {
		fmt.Printf("%s의 나이는 %d\n", name, age)
	}
	/*
		lines, err := datefile.Getstrings("votes.txt")
		if err != nil {
			log.Fatal(err)
		}
		var names []string
		var counts []int
		for _, line := range lines {
			matched := false
			for i, name := range names {
				if name == line {
					counts[i]++
					matched = true
				}
			}
			if matched == false {
				names = append(names, line)
				counts = append(counts, 1)
			}
		}
		for i, name := range names {
			fmt.Printf("%s: %d\n", name, counts[i])
		}*/
}
