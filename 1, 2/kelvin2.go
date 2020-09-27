package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	k -= 273.15
	return celsius(k)
}

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

func kelvinToFahrenheit(k kelvin) fahrenheit {
	c := kelvinToCelsius(k)
	return fahrenheit(celsiusToFahrenheit(c))
}

func main() {
	var k kelvin = 233.0
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	fmt.Print(k, "° K is ", c, "° C\n")
	fmt.Print(c, "° C is ", f, " Fahrenheit\n")
	//fmt.Print(kelvin, "° K is ", fahrenheit, " Fahrenheit\n")
	fmt.Printf("0° K is %.2f° F\n", kelvinToFahrenheit(0.0))
}
