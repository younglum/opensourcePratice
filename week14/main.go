package main

import "fmt"

/*
	type student struct {
		name  string
		grade float32
		id    int
	}
*/
type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	//visitors : 구조체 슬라이스
	total_cost := 0
	for _, v := range visitors {
		total_cost += v.cost
	}
	return total_cost
}

func main() {
	var num_visitors int
	fmt.Println("How many visitors?")
	fmt.Scanln(&num_visitors)
	vs := make([]visitor, num_visitors)

	for i := 0; i < num_visitors; i++ {
		var age int
		fmt.Printf("%d번째 사람의 나이를 입력하세요: ", i+1)
		fmt.Scan(&age)
		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 9000}
		}
	}
	fmt.Printf("전체 금액은 %d원 입니다", calculateCost(vs))

	/*
		var student1 student

		student1.id = 202444076
		student1.name = "임수영"
		student1.grade = 4.5
		fmt.Println(student1.grade)
		var student2 student

		student2.id = 202444123
		student2.name = "홍길동"
		student2.grade = 4.2
		fmt.Println(student2.id)

			ages := make(map[string]int)
			var name string
			var age int
			ages_slice := []string{}
			for {
				fmt.Print("이름을 입력하세요 : ")
				fmt.Scanln(&name)
				if name == "q" || name == "Q" {
					break
				}
				fmt.Print("나이를 입력하세요 : ")
				fmt.Scanln(&age)
				ages[name] = age
				ages_slice = append(ages_slice, ages[name])
			}
			for name, age := range ages_slice {
				fmt.Printf("%s의 나이는 %d\n", name, age)
			}
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
