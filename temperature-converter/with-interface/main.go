package main

import "fmt"

type celcius struct {
	temperature float64
}

type fahrenheit struct {
	temperature float64
}

type kelvin struct {
	temperature float64
}

func (c celcius) toCelcius() float64 {
	return c.temperature
}

func (c celcius) toFahrenheit() float64 {
	return (9.0 / 5.0 * c.temperature) + 32
}

func (c celcius) toKelvin() float64 {
	return c.temperature + 273.15
}

func (f fahrenheit) toCelcius() float64 {
	return (5.0 / 9.0) * (f.temperature - 32)
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.temperature
}

func (f fahrenheit) toKelvin() float64 {
	return (f.temperature - 459.67) * (5.0 / 9.0)
}

func (k kelvin) toCelcius() float64 {
	return k.temperature - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.temperature * (9.0 / 5.0)) - 459.67
}

func (k kelvin) toKelvin() float64 {
	return k.temperature
}

type calculateTemperature interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Select a first temperature unit")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")

	var firstUnit int
	fmt.Print("Enter the first temperature unit: ")
	fmt.Scan(&firstUnit)

	for firstUnit < 1 || firstUnit > 3 {
		fmt.Println("Invalid temperature unit")
		fmt.Print("Enter the first temperature unit: ")
		fmt.Scan(&firstUnit)
	}

	fmt.Println("Select a result temperature unit")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")

	var resultUnit int
	fmt.Print("Enter the result temperature unit: ")
	fmt.Scan(&resultUnit)

	for resultUnit < 1 || resultUnit > 3 {
		fmt.Println("Invalid temperature unit")
		fmt.Print("Enter the result temperature unit: ")
		fmt.Scan(&resultUnit)
	}

	var temperature float64
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temperature)

	var interfaceTemperature calculateTemperature
	switch firstUnit {
	case 1:
		interfaceTemperature = celcius{temperature}
	case 2:
		interfaceTemperature = fahrenheit{temperature}
	case 3:
		interfaceTemperature = kelvin{temperature}
	}

	var resultTemperature float64
	switch resultUnit {
	case 1:
		resultTemperature = interfaceTemperature.toCelcius()
	case 2:
		resultTemperature = interfaceTemperature.toFahrenheit()
	case 3:
		resultTemperature = interfaceTemperature.toKelvin()
	}
	fmt.Printf("Result temperature: %.2f\n", resultTemperature)
}
