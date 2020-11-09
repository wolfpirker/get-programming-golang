package main

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func sortFoods(f []string, less func(i, j int) bool) {
	// sort by the length of the food; number of letters upwards
	if less == nil {
		less = func(i, j int) bool { return len(f[i]) < len(f[j]) }
	}
	sort.Slice(f, less)
}

func main() {
	food := []string{"onion", "carrot", "brocolli", "celery"}
	sortFoods(food, nil)
	fmt.Println(food)
}
