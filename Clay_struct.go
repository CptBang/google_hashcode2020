package main

import (
	"github.com/edwingeng/deque"
	"github.com/emirpasic/gods/maps/treemap"
	// "github.com/emirpasic/gods/utils"
)

// "github.com/emirpasic/gods/maps/treemap"

// Elem var dq deque.Deque
// func (*Library) compare(a, b Library) int {
// 	aAsserted := a.rating.(int)
// 	bAsserted := b.rating.(int)
// 	switch {
// 	case aAsserted > bAsserted:
// 		return 1
// 	case aAsserted < bAsserted:
// 		return -1
// 	default:
// 		return 0
// 	}
// }

// func (l *Library) getRate() int {
// 	return l.rating
// }

// type libElem interface {
// 	getRate() int
// 	compare(a, b Library) int
// }

// type libComp func(a, b libElem) int

// func compare(a libElem, b libElem) int {
// 	return a.getRate() - b.getRate()
// }

// var state = treemap.NewWith(utils.Comparator)
var state = treemap.NewWithIntComparator()

// state := make([][]int, )

// build catalog
// func addLib(inp LibElem) {
func addLib(inp *Library) {
	// var val libElem
	// var found bool
	_, found := state.Get(inp.rating)
	// val := v.(*Library)
	// val := v.(deque.Deque)

	if found == true {
		state.Put(inp.rating, inp)
		// if val != nil {
		// 	// val = append(val, inp)
		// 	// val = deque.Deque(val)
		// 	val.PushBack(inp)
		// } else {
		// 	state.Put(inp.rating, inp)
		// 	// val = deqeue.NewDeque()
		// 	// val.PushBack(inp)
		// }
	} else {
		state.Put(inp.rating, inp)
	}
}

// add library to treemap

// remove
func delLib(inp *Library) {
	// v, found := state.Get(inp.rating)
	state.Remove(inp.rating)
}

// get top library
func getBest() *Library {
	// var key int
	// var val *Library
	// var val libElem
	_, v := state.Max()
	val := v.(deque.Deque)
	b := val.Peek(0)
	best := b.(*Library)
	return best

	// return val
}
