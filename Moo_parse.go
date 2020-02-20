package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// func createLibrary(lib1 string, lib2 string) Library struct {

// 	NewLibrary()

// }

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

	line2 := strings.Fields(result[1])
	// bookScores :=

	// create all libraries and put into array
	// var libraries [numLibraries]Library
	// for i := 0; i < len(libraries); i++{
	// 	for j := 2; j <
	// 	createLibrary()
	// }

	fmt.Printf("%s", line1)
	fmt.Printf("%s", line2)
	fmt.Printf("%d", numBooks)
	fmt.Printf("%d", numLibraries)
	fmt.Printf("%d", numDays)

	// for i := 0; i < K; i++ {
	// 	intPizzas[i], _ = strconv.Atoi(strPizzas[i])
	// }
}

// array
//
//
