package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Split(text, " ")
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.TrimRight(word, ",")
		word = strings.TrimRight(word, ".")
	}

	for _, t := range words {
		frequency[t]++
	}
	return frequency
}

func main() {
	text :=
		`Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.
		 Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, 
		 nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, 
		 sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
		  arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. 
		  Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula,
		   porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus.
			Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue.
			 Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus,
			  sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar,
			   hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. 
		Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. 
`
	frequency := countWords(text)

	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}

}
