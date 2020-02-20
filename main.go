package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	parseFile(file)
	// daysLeft = 7
	// bookScores = []int{1,2,3,6,5,4}
	// addLib(makeLibrary(0, 5, 2, 2, []int{0,1,2,3,4}))
	// addLib(makeLibrary(1, 4, 3, 1, []int{3,2,5,0}))
	// addLib(makeLibrary(2, bookCount, signupLen, shipRate, parseBookIds(secondLine)))
	doneList := createRatedList()
	submit(doneList)

	fmt.Printf("doneList(%v)\n", doneList)


	// solution := Clayton()
	// fmt.Println(solution)
}
