package main

import "fmt"

type character struct {
	name string
	leftHand *item
}

type item struct {
	name string
	valid bool
}

func (c *character) pickup(i *item){
	if c != nil && i.valid {
		fmt.Println(c.name, " picks up a ", i.name)
		c.leftHand = i
	} else {
		return
	}	
}

func (c *character) give(to *character){
	if c != nil && c.leftHand != nil && to.leftHand == nil {
		fmt.Println(c.name, " gives a ", c.leftHand, " to ", to.name)
		to.leftHand = c.leftHand
		c.leftHand = nil
	} else if (c.leftHand == nil) {
		fmt.Println(c.name, " has nothing to give!")		
	} else if (to.leftHand != nil) {
		fmt.Println(to.name, " has an item already!")		
	}
	

	
}

func newItem(n string) item {
	return item{name: n, valid: true}
}

func (i item) String() string {
	if !i.valid {
		return "not valid"
	}
	return fmt.Sprintf("%s", i.name)
}

func (c character) String() string {
	return c.name
}

func main() {
	hero := &character{name: "Arthur"}
	knight := &character{name: "Gwydyan"}
	fmt.Println("Hero ", hero)
	fmt.Println("Knight ", knight)
	
	hg := newItem("Holy Grail")
	//hero.leftHand = &hg
	hero.pickup(&hg)
	hero.give(knight)

	

}
