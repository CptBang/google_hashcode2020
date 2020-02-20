package main

import (
    "fmt"
    "os"
)

func check(err error) {
	if err != nil {
		panic(err)
    }
}

func submit(libraries []*Library) {
	// file := os.Args[1]
	submission, _ := os.Create("submit")

	defer submission.Close()

	_, err := submission.WriteString(fmt.Sprintf("%d\n", len(libraries)))
	check(err)

	for _, library := range libraries{
		_, err := submission.WriteString(fmt.Sprintf("%d %d\n", library.libID, library.bookCount))
		check(err)
		for _, id := range library.booksIDs {
			_, err := submission.WriteString(fmt.Sprintf("%d", id))
			check(err)
		}
	}
}
