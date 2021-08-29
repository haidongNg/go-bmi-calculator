package main

import (
	"runtime"

	"github.com/haidongNg/go-bmi-calculator/info"
)

func main() {
	// Output infomation
	info.PrintWelCome()

	// Prompt for user input (weight + height)

	// Save that user input in variables
	currentRuntime := "\n"

	if runtime.GOOS == "windows" {
		currentRuntime = "\r\n"
	}

	weight, height := getUserMetrics(currentRuntime)

	// Caculate the BMI (weight (height * height))
	bmi := calculateBIM(weight, height)

	// Output the calculated BMI
	printBMI(bmi)
}

func calculateBIM(weight float64, height float64) float64 {
	return weight / (height * height)
}
