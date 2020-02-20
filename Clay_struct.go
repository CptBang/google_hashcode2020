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
	v, found := state.Get(inp.rating)
	// val := v.(*Library)

	if found || v != nil {
		state.Put(inp.rating, inp)
		val := v.(deque.Deque)
		if val != nil {
			val.PushFront(inp)
		} else {
			n := deque.NewDeque()
			inp.rate()
			n.PushFront(inp)
			state.Put(inp.rating, n)
		}
	} else {
		n := deque.NewDeque()
		inp.rate()
		n.PushFront(inp)
		state.Put(inp.rating, n)
	}
}

// add library to treemap

// rerate get base and call rate if different remove and reinsert
func rerateLib() {
	chk := getBest()
	prev := chk.rating
	chk.rate()
	if chk.rating != prev {
		delLib(chk)
		addLib(chk)
	}
}

// remove
func delLib(inp *Library) {
	v, found := state.Get(inp.rating)
	val := v.(deque.Deque)
	if found {
		if val.Empty() {
			state.Remove(inp.rating)
		} else {
			val.PopFront()
		}
	}
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
