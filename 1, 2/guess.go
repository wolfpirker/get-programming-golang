package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const my_num = 12
	var i = 0
	var guess = 0
	for {
		guess = rand.Intn(100) + 1
		i++
		if guess < my_num {
			fmt.Printf("goes guess num %v was smaller\n", i)
		} else if guess > my_num {
			fmt.Printf("goes guess num %v was bigger\n", i)
		} else if guess == my_num {
			fmt.Println("go guess the right number!")
			fmt.Printf("go needed %v guesses\n", i)
			break
		}
	}
}
