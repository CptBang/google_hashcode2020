package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// order as many slices as possible, but not more than M (first value in file) and none of the identical pizza

// scan lines of file until space
//		first two are always M - max number of slices, K - different types of pizza
//		rest are the number of slices of each pizza stored in separate structure

// dynamic programming to remember the numbers of slices
// recursive?

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

// dynamic array
// at the dp we save the closest to M value
// if it's greater than slices then we say it's equal to slices
// loop through
//	if number of slices is greater than previous and less than or equal to M, then set slices
//
func recurSlices(M int, K int, intPizzas []int, slices int, value int) int {
	if slices >= M || value >= len(intPizzas) {
		return 0
	}

	sliceOne := intPizzas[value] + recurSlices(M, K, intPizzas, slices, value+1)
	sliceTwo := intPizzas[value] + recurSlices(M, K, intPizzas, slices, value+2)

	if sliceOne >= sliceTwo && sliceOne+slices < M {
		slices += sliceOne
	} else if sliceTwo > sliceOne && sliceTwo+slices < M {
		slices += sliceTwo
	}
	return slices
}

func getSlices(M int, K int, intPizzas []int) int {
	// solution := make([]string, 2)
	slices := recurSlices(M, K, intPizzas, 0, 0)
	return slices
}

func main() {
	file := os.Args[1]
	result := scanText(file)
	MK := strings.Fields(result[0])
	strPizzas := strings.Fields(result[1])
	M, _ := strconv.Atoi(MK[0]) //total slices
	K, _ := strconv.Atoi(MK[1]) //number of pizzas
	intPizzas := make([]int, K+1)
	for i := 0; i < K; i++ {
		intPizzas[i], _ = strconv.Atoi(strPizzas[i])
	}
	solution := getSlices(M, K, intPizzas)
	// fmt.Println(solution[0])
	// fmt.Println(solution[1])
	fmt.Println(solution)
}
