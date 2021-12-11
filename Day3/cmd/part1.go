package main

import (
	"bufio"
	"fmt"
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
	var gamma uint64 = 0
	var epsilon uint64 = 0
	var curNum uint64 = 0
	counter := 0
	dataSlice := make([]uint64, 0, 0)
	filePath := argsWithoutProg[0]

	// Open the file given as an argument
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Scan the Values and convert them to strings and store in a slice
	for scanner.Scan() {
		curNum, err = strconv.ParseUint(scanner.Text(), 2, 16)
		check(err)
		dataSlice = append(dataSlice, curNum)
	}

	for i := 11; i >= 0; i-- {
		for _, j := range dataSlice {
			num := j << (64 - i - 1)
			num >>= 63
			if num != 0 {
				counter++
			}
		}
		gamma <<= 1
		if counter > len(dataSlice)/2 {
			gamma += 1
		}
		counter = 0
	}

	epsilon = (^gamma)
	epsilon <<= 52
	epsilon >>= 52

	fmt.Printf("Gamma Rate:        %d\n", gamma)
	fmt.Printf("Epsilon Rate:      %d\n", epsilon)
	fmt.Printf("Power Consumption: %d\n", gamma*epsilon)
}
