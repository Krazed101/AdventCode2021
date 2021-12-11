package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	depth := 0
	horizontal := 0
	aim := 0
	curNum := 0
	dataSlice := make([]string, 0, 0)
	filePath := argsWithoutProg[0]

	// Open the file given as an argument
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Scan the Values and convert them to strings and store in a slice
	for scanner.Scan() {
		dataSlice = append(dataSlice, scanner.Text())
		check(err)
	}

	for _, i := range dataSlice {
		s := strings.Split(i, " ")
		curNum, err = strconv.Atoi(s[1])
		check(err)
		switch s[0] {
		case "forward":
			horizontal += curNum
			depth += aim * curNum
		case "down":
			aim += curNum
		case "up":
			aim -= curNum
		}

	}

	fmt.Printf("Horizontal Movement: %d\n", horizontal)
	fmt.Printf("Depth Movement:      %d\n", depth)
	fmt.Printf("Total Distance:      %d\n", horizontal*depth)
}
