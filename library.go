package main

import "fmt"

//Library is a struct for holding library info for book scanning
type Library struct {
	libID int
	bookCount int
	signupLen int
	shipRate int
	booksIDs []int

	rating int
}

func makeLibrary(id, bookCount, signupLen, shipRate int, bookIDs []int) *Library {
	l := new(Library)
	l.libID = id
	l.bookCount = bookCount
	l.signupLen = signupLen
	l.shipRate = shipRate
	l.booksIDs = bookIDs

	l.rate();

	return l
}

func (l *Library) rate() {
	timeLen := l.signupLen + (l.bookCount / l.shipRate)
	totalScore := 0
	for i := 0; i < l.bookCount; i++ {
		totalScore += bookScores[l.booksIDs[i]]
	}

	l.rating = totalScore / timeLen
}

func (l *Library) String() string {
	// return fmt.Sprintf("<%d>  %d books, %d books/day, %d signup.  \t Rating:%d\tids: %v\n", l.libID, l.bookCount, l.shipRate, l.signupLen, l.rating, l.booksIDs)
	return fmt.Sprintf("<%d>  %d books, %d books/day, %d signup.  \t Rating:%d\n", l.libID, l.bookCount, l.shipRate, l.signupLen, l.rating)
}
