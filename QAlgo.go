package main

func createRatedList() []*Library {
	var libList []*Library


	for daysLeft > 0 {
		var currLib *Library
		for currLib != getBest() {
			currLib = getBest()
			rerateLib()
		}
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