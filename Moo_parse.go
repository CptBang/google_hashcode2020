package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
)

var bookScores []int
var daysLeft int

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

func parseFile(ofile string) {
	// result := scanText(file)

	numBooks, numLibraries, numDays := 0, 0, 0

	file, err := os.Open(ofile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	read := bufio.NewReader(file)


		ln1, _ := read.ReadString('\n')
		line1 := strings.Fields(ln1)
		numBooks, _ = strconv.Atoi(line1[0]) // total number of books
		numLibraries, _ = strconv.Atoi(line1[1]) // total number of libraries
		numDays, _ = strconv.Atoi(line1[2]) // total number of days for scanning
		daysLeft = numDays

		ln2, _ := read.ReadString('\n')
		line2 := strings.Fields(ln2)
		bookScores = make([]int, numBooks)
		for i := 0; i < numBooks; i++ {
			bookScores[i], _ = strconv.Atoi(line2[i])
		}

		//create all libraries and put into array
		for i := 0; i < numLibraries; i++ {

			ln3, _ := read.ReadString('\n')
			firstLine := strings.Fields(ln3)

			ln4, _ := read.ReadString('\n')
			secondLine := strings.Fields(ln4)

			bookCount, _ := strconv.Atoi(firstLine[0])
			signupLen, _ := strconv.Atoi(firstLine[1])
			shipRate, _ := strconv.Atoi(firstLine[2])
			addLib(makeLibrary(i, bookCount, signupLen, shipRate, parseBookIds(secondLine)))


	}
	// fmt.Printf("%d %d %d\n", numLibraries, numBooks, numDays)
	// fmt.Printf("%v\n", bookScores)
}

