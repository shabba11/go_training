package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originalTemp := createTemperatureAdjuster()
	fmt.Printf("Original temperature is %.1f\n", originalTemp)

	fmt.Printf("Adjusted temp +1.5: %.1f grad C\n", adjustTemp(1.5))  // 91.5
	fmt.Printf("Adjusted temp -3.0: %.1f grad C\n", adjustTemp(-3.0)) // 88.5
	fmt.Printf("Adjusted temp +5.0: %.1f grad C\n", adjustTemp(5.0))  // 93.5

	fmt.Printf("Original temperature is %.1f\n", originalTemp) // 90.0
}
