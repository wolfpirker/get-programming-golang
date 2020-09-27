package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(15) + 12)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var k kelvin = 5.5
	sensor := calibrate(realSensor, k)
	fmt.Println(sensor())

	/*
		k = 7.5
		fmt.Println(sensor())
		sensor = calibrate(realSensor, k)
		fmt.Println(sensor())
	*/

	sensor = calibrate(fakeSensor, k)
	fmt.Println(sensor())
	fmt.Println(sensor())
}
