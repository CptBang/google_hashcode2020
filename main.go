package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := os.Args[1]

	solution := getSlices(M, K, intPizzas)
	fmt.Println(solution)
}
