package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	counter := 0
	curNum := 0
	dataSlice := make([]int, 0, 0)
	prevTotal := math.MaxInt
	curTotal := 0
	filePath := argsWithoutProg[0]

	// Open the file given as an argument
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Scan the Values and convert them to integers and store in a slice
	for scanner.Scan() {
		curNum, err = strconv.Atoi(scanner.Text())
		check(err)
		dataSlice = append(dataSlice, curNum)
	}

	for i := 0; i < len(dataSlice)-2; i++ {
		curTotal = dataSlice[i] + dataSlice[i+1] + dataSlice[i+2]
		if curTotal > prevTotal {
			counter++
		}
		prevTotal = curTotal
	}

	fmt.Printf("Number of Increases: %d\n", counter)
}
