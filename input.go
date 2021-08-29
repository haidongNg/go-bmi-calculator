package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/haidongNg/go-bmi-calculator/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics(currentRuntime string) (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt, currentRuntime)
	height = getUserInput(info.HeightPrompt, currentRuntime)

	return
}

func getUserInput(promptText string, currentRuntime string) (value float64) {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, currentRuntime, "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
