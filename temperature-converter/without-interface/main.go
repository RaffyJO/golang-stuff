package main

import "fmt"

func main() {
	fmt.Println("Select a temperature unit:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Celsius to Kelvin")
	fmt.Println("3. Fahrenheit to Celsius")
	fmt.Println("4. Fahrenheit to Kelvin")
	fmt.Println("5. Kelvin to Celsius")
	fmt.Println("6. Kelvin to Fahrenheit")

	var unit int
	fmt.Print("Enter the temperature unit: ")
	fmt.Scan(&unit)

	for unit < 1 || unit > 6 {
		fmt.Println("Invalid temperature unit")
		fmt.Print("Enter the temperature unit: ")
		fmt.Scan(&unit)
	}

	var temperature float64
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temperature)

	switch unit {
	case 1:
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temperature, CelsiusToFahrenheit(temperature))
	case 2:
		fmt.Printf("%.2f Celsius is %.2f Kelvin\n", temperature, CelsiusToKelvin(temperature))
	case 3:
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temperature, FahrenheitToCelsius(temperature))
	case 4:
		fmt.Printf("%.2f Fahrenheit is %.2f Kelvin\n", temperature, FahrenheitToKelvin(temperature))
	case 5:
		fmt.Printf("%.2f Kelvin is %.2f Celsius\n", temperature, KelvinToCelsius(temperature))
	default:
		fmt.Printf("%.2f Kelvin is %.2f Fahrenheit\n", temperature, KelvinToFahrenheit(temperature))
	}
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (9.0 / 5.0 * celsius) + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (5.0 / 9.0) * (fahrenheit - 32)
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit - 459.67) * (5.0 / 9.0)
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin * (9.0 / 5.0)) - 459.67
}
