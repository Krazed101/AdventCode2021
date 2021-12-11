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
	counter := 0
	curNum := 0
	filePath := argsWithoutProg[0]

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prevNum, err := strconv.Atoi(scanner.Text())
	check(err)
	for scanner.Scan() {
		curNum, err = strconv.Atoi(scanner.Text())
		check(err)
		if curNum > prevNum {
			counter++
		}
		prevNum = curNum
	}
	fmt.Printf("Number of Increases: %d\n", counter)

}
