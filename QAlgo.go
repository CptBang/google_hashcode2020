package main

import (
)

func createRatedList() []*Library {
	var libList []*Library


	for daysLeft > 0 {
		// fmt.Printf("state.Size( %d )\n", state.Size())
		// var currLib *Library
		// for currLib != getBest() {
		// 	fmt.Printf("getBest first: %s\n", currLib)
		// 	currLib = getBest()
		// 	fmt.Printf("getBest second: %s\n", currLib)
		// 	rerateLib()
		// }
		// fmt.Printf("state.Size( %d )\n", state.Size())
		// fmt.Printf("currLib: %s", currLib)

		currLib := getBest()
		if (currLib == nil) {
			break
		}
		delLib(currLib)
		currLib.sortIDsByScore()
		libList = append(libList, currLib)
		for i := 0; i < currLib.shipRate * (daysLeft - currLib.signupLen); i++ {
			if (i >= currLib.bookCount) {
				break
			}
			bookScores[currLib.booksIDs[i]] = 0
		}
		daysLeft -= currLib.signupLen
		rerateLib()
	}

	return libList
}

//21,725,028 original
//