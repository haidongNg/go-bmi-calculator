package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/haidongNg/go-bmi-calculator/info"
)

func main() {
	fmt.Println("Hello word!")

	// Output infomation
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	// Prompt for user input (weight + height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	currentRuntime := "\n"

	if runtime.GOOS == "windows" {
		currentRuntime = "\r\n"
	}

	weightInput = strings.Replace(weightInput, currentRuntime, "", -1)
	heightInput = strings.Replace(heightInput, currentRuntime, "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Caculate the BMI (weight (height * height))
	bmi := (weight / (height * height)) * 703

	// Output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
