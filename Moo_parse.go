package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var bookScores []int
var numDays int

func parseBookIds(secondLine []string) []int {
	bookIds := make([]int, len(secondLine))
	for i := 0; i < len(secondLine); i++ {
		bookIds[i], _ = strconv.Atoi(secondLine[i])
	}

	return bookIds
}

func scanText(file string) []string {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func parseFile(file string) {
	result := scanText(file)

	line1 := strings.Fields(result[0])
	numBooks, _ := strconv.Atoi(line1[0]) // total number of books
	numLibraries, _ := strconv.Atoi(line1[1]) // total number of libraries
	numDays, _ := strconv.Atoi(line1[2]) // total number of days for scanning

	// create array of book scores
	line2 := strings.Fields(result[1])
	bookScores = make([]int, numBooks)
	for i := 0; i < numBooks; i++ {
		bookScores[i], _ = strconv.Atoi(line2[i])
	}

	//create all libraries and put into array
	for i := 0; i < numLibraries; i++{
			lineNum := 2+(i*2)
			firstLine := strings.Fields(result[lineNum])

			secondLine := strings.Fields(result[lineNum+1])

			bookCount, _ := strconv.Atoi(firstLine[0])
			signupLen, _ := strconv.Atoi(firstLine[1])
			shipRate, _ := strconv.Atoi(firstLine[2])
			addLib(makeLibrary(i, bookCount, signupLen, shipRate, parseBookIds(secondLine)))
	}

	
	fmt.Printf("%d %d %d\n", numLibraries, numBooks, numDays)
	fmt.Printf("%v\n", bookScores)
}

// array
//
//
