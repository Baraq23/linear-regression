package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathskills/stats"
)

func main() {
	// checking if input is correctly entered
	arguments := len(os.Args)
	if arguments != 2 {
		fmt.Println("Usage: go run . data.txt")
		return
	} else if os.Args[1] != "data.txt" {
		fmt.Println("Input should be: go run program-name.go data.txt")
		return
	}
	// open and read data file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: Data file not found. Please provide data.txt")
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numsY := []float64{}
	numsX := []float64{}
	lineNum := 0
	for scanner.Scan() {
		str := scanner.Text()
		// Removing trailing or leading spaces
		str = strings.TrimSpace(str)
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("Error: Line %v is an invalid data\n", lineNum)
			os.Exit(0)
		}
		// Number bigger than the maximum float64 cappacity will not be computed
		if num > math.MaxInt64 {
			fmt.Printf("Error: Data at line %v in data file is too big to compute!\n", lineNum)
			os.Exit(0)
		}
		numsX = append(numsX, float64(lineNum))
		numsY = append(numsY, float64(num))
		lineNum++
	}
	// this check if the data file has values
	if len(numsY) == 0 {
		fmt.Println("Error: Empty data file found.")
		os.Exit(0)
	}
	b, a, r := stats.LinearStats(numsX, numsY)

	// b & a to 6 decimal places
	b = (math.Round(b * 1000000)) / 1000000
	a = (math.Round(a * 1000000)) / 1000000
	// r to 10 decimal places
	r = (math.Round(r * 10000000000)) / 10000000000

	// Printing results
	if a >= 0 {
		fmt.Printf("Linear Regression line: y = %.6fx + %.6f\n", b, a)
	} else {
		fmt.Printf("Linear Regression line: y = %.6fx %.6f\n", b, a)
	}
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
