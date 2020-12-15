package main

import (
	"flag" // command-line flag parsing
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"time"
)

type problem struct {
	q string // question
	a string // answer
}

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")

	// see documentation https://golang.org/pkg/time/#Timer
	// ticker similar to timer, but timer fires just once
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	

	file, err := os.Open(*csvFilename)
	if err != nil{
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("Failed to read file: %s", file))
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)	

	correct := 0
	problemloop: // label
	for i, p := range problems {		
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string			
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case  <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}