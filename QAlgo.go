package main

import (
)

func createRatedList() []*Library {
	var libList []*Library


	for daysLeft > 0 {
		// fmt.Printf("state.Size( %d )\n", state.Size())
		// var currLib *Library
		currLib := getBest()
		// for currLib != getBest() {
		// 	fmt.Printf("getBest first: %s\n", currLib)
		// 	currLib = getBest()
		// 	fmt.Printf("getBest second: %s\n", currLib)
		// 	rerateLib()
		// }
		// fmt.Printf("state.Size( %d )\n", state.Size())
		// fmt.Printf("currLib: %s", currLib)
		if (currLib == nil) {
			break
		}
		delLib(currLib)
		//currLib.sortIDsByScore()
		libList = append(libList, currLib)
		for _, v := range currLib.booksIDs {
			bookScores[v] = 0;
		}
		daysLeft -= currLib.signupLen
	}

	return libList
}
