package main

// Youtube Turoial: 
// Tensor Programming: Building a Basic RPC Server and Client with Go

import (
	"fmt"
)


type Item struct {
	title string
	body string
}

var database []Item

func GetByName(title string) Item{
	var getItem Item
	
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteItem(item Item) Item {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			// create new database without the item we want to be removed
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}

	return del
}
func main() {
	fmt.Println("initial database: ", database)
	a := Item{"1st", "a test item"}
	b := Item{"2nd", "a second item"}
	c := Item{"3rd", "a third item"}
	
	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("2nd database: ", database)

	DeleteItem(b)
	fmt.Println("3rd database: ", database)

	EditItem("3rd", Item{"fourth", "a new item"})
	fmt.Println("4th database: ", database)

	x := GetByName("4th")
	y := GetByName("1st")
	fmt.Println(x, y)
}