package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"strings"
)

type problem struct {
	q string // question
	a string // answer
}

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")
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

	correct := 0
	
	for i, p := range problems {
		// could be broke up into another function, minute 17 part 1
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You score %d out of %d\n", correct, len(problems))
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