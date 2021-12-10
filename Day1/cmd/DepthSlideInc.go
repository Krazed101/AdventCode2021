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
	prevTotal := math.MaxInt
	curTotal := 0
	i := 4
	scanA := [3]int{0, 0, 0}
	scanB := [3]int{0, 0, 0}
	scanC := [3]int{0, 0, 0}
	scanD := [3]int{0, 0, 0}
	filePath := argsWithoutProg[0]

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanA[0], err = strconv.Atoi(scanner.Text())
	check(err)
	scanA[1], err = strconv.Atoi(scanner.Text())
	check(err)
	scanA[2], err = strconv.Atoi(scanner.Text())
	check(err)

	scanB[0] = scanA[1]
	scanB[1] = scanA[2]

	scanC[0] = scanA[2]

	for scanner.Scan() {
		curNum, err = strconv.Atoi(scanner.Text())
		check(err)
		switch i % 4 {
		case 0:
			scanB[2] = curNum
			scanC[1] = curNum
			scanD[0] = curNum

			curTotal = scanA[0] + scanA[1] + scanA[2]

		case 1:
			scanA[0] = curNum
			scanC[2] = curNum
			scanD[1] = curNum

			curTotal = scanB[0] + scanB[1] + scanB[2]

		case 2:
			scanA[1] = curNum
			scanB[0] = curNum
			scanD[2] = curNum

			curTotal = scanC[0] + scanC[1] + scanC[2]

		case 3:
			scanA[2] = curNum
			scanB[1] = curNum
			scanC[0] = curNum

			curTotal = scanD[0] + scanD[1] + scanD[2]
		}

		i++
		/* Debug code
		fmt.Printf("PrevTotal: %d\n", prevTotal)
		fmt.Printf("CurrTotal: %d\n", curTotal)
		fmt.Printf("Counter  : %d\n", counter)

		fmt.Scanln()*/

		if curTotal > prevTotal {
			counter++
		}
		prevTotal = curTotal
	}
	fmt.Printf("Number of Increases: %d\n", counter)

}
