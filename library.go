package main

//Library is a struct for holding library info for book scanning
type Library struct {
	bookCount int
	signupLen int
	shipRate int
	booksIDs []int

	rating int
}

//NewLibrary constructor
func NewLibrary(bookCount, signupLen, shipRate int, bookIDs []int) *Library {
	l := new(Library)
	l.bookCount = bookCount
	l.signupLen = signupLen
	l.shipRate = shipRate
	l.booksIDs = bookIDs

	l.rate();

	return l
}

func (l *Library) rate() {
	l.rating = 0
}
