package main

import "fmt"

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return ((c * 9.0 / 5.0) + 32)
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celsiusToFahrenheit(c)
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Print(kelvin, "° K is ", celsius, "° C\n")
	fmt.Print(celsius, "° C is ", fahrenheit, " Fahrenheit\n")
	//fmt.Print(kelvin, "° K is ", fahrenheit, " Fahrenheit\n")
	fmt.Printf("0° K is %.2f° F\n", kelvinToFahrenheit(0.0))
}
