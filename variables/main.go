package main

import "fmt"

const oneFeetToMeter float64 = 0.3048

func main() {
	fmt.Print("Enter temprature in fahrenheit: ")
	var tempratureInFahrenheit float64
	fmt.Scanf("%f", &tempratureInFahrenheit)

	fmt.Print("Enter distance in feet: ")
	var distanceInFeet float64
	fmt.Scanf("%f", &distanceInFeet)

	temperatureInCelsius := (tempratureInFahrenheit - 32) * 5 / 9
	distanceInMeters := oneFeetToMeter * distanceInFeet

	fmt.Println(tempratureInFahrenheit, "degrees Fahrenheit = ", temperatureInCelsius, "degrees Celsius")
	fmt.Println(distanceInFeet, "feet = ", distanceInMeters, "meters")

}

/*
	1) the 2 ways 0f defining variable is by strict typing or interference
		var x string = "Hello",
	 	x := "Hello"
		x := 5; x += 1?
	2) The value of x
		x := 5; x += 1?
		x = 6
	3) scope of a variable is where the variable can be used. The scope of a variable is in Go is determined by the block in which the vriable is in
	that is the closest open an closed brackets {}

	4) The value of const can't change once it has been initialized but the value of var can change

*/
