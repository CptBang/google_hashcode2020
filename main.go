package main

import (
	"os"
)

func main() {
	file := os.Args[1]
	parseFile(file)
	doneList := createRatedList()
	submit(doneList)

	// fmt.Printf("%v", doneList)

	// solution := Clayton()
	// fmt.Println(solution)
}
