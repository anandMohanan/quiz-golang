package main

import (
	"fmt"
	"time"
)

func gq(lines [][]string) {
	problems := parseLines(lines)
	//fmt.Println(problems)
	correct := 0

	Timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem #%d , %s = ", i+1, p.Question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()
		select {
		case <-Timer.C:
			fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.Answer {
				fmt.Println("correct answer")
				correct++
			} else {
				fmt.Println("wrong answer")
			}

		}

	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}
