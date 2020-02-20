package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	parseFile(file)
	doneList := createRatedList()

	fmt.Printf("%v", doneList)

	
	// solution := Clayton()
	// fmt.Println(solution)
}
