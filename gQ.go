package main

import "fmt"

func gq(lines [][]string) {
	problems := parseLines(lines)
	//fmt.Println(problems)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d , %s = \n", i+1, p.Question)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == p.Answer {
			fmt.Println("correct answer")
			correct++
		} else {
			fmt.Println("wrong answer")
		}
	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}
