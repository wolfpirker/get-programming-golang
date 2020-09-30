package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "====================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

// function to get data for each row -> why not just an integer variable?
type getRowFn func(row int) (string, string)

// drawTable draws a two column table
func drawTable(header1 string, header2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Println(rowFormat, header1, header2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func cToF(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func main() {
	drawTable("°C", "F", 20, cToF)
	//drawTable("°C", "F", 20, fToC)
}
