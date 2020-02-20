package main

import (
	// "github.com/edwingeng/deque"
	// "github.com/emirpasic/gods/maps/treemap"
	// "fmt"
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
// var state = treemap.NewWithIntComparator()
var libs []*Library

// state := make([][]int, )

// build catalog
// func addLib(inp LibElem) {
func addLib(inp *Library) {
	// // var val libElem
	// // var found bool
	// fmt.Printf("addLib( %s )\n", inp)
	// v, found := state.Get(inp.rating)
	// // val := v.(*Library)
	// if found && v != nil {
	// 	// state.Put(inp.rating, inp)
	// 	val := v.(deque.Deque)
	// 	if val == nil {
	// 		panic("val nil in addLib")
	// 	}
	// 	fmt.Printf("\tpush front addLib( %s )\n", inp)
	// 	val.PushFront(inp)
	// } else {
	// 	n := deque.NewDeque()
	// 	// inp.rate()
	// 	n.PushFront(inp)
	// 	fmt.Printf("\tNewDeque( %s )\n", n.Front())
	// 	state.Put(inp.rating, n)
	// }
	libs = append(libs, inp)
}

// add library to treemap

// rerate get base and call rate if different remove and reinsert
func rerateLib() {
	for _,lib := range libs {
		lib.rate()
	}
	// chk := getBest()
	// prev := chk.rating
	// chk.rate()
	// if chk.rating != prev {
	// 	delLib(chk)
	// 	addLib(chk)
	// }
}

// remove
func delLib(inp *Library) {
	for	i,v := range libs {
		if v == inp {
			libs = append(libs[:i], libs[i+1:]...)
		}
	}
	// if inp == nil {
	// 	panic("inp nil in delLib")
	// }
	// v, found := state.Get(inp.rating)
	// if found {
	// 	if v == nil {
	// 		panic("v is nil in delLib")
	// 	}
	// 	val := v.(deque.Deque)
	// 	// if val.Empty() {
	// 	// 	// state.Remove(inp.rating)
	// 	// } 
	// 	if val.Empty() == false {
	// 		val.PopFront()
	// 	} else {
	// 		fmt.Println("delLib empty val")
	// 	}
	// }

}

// TODO: sort bookScores and keep bookIDs synced

// get top library
func getBest() *Library {
	max := 0
	var out *Library
	for _,lib := range libs {
		if max < lib.rating {
			max = lib.rating
			out = lib
		}
	}
	// var key int
	// var val *Library
	// var val libElem
	// _, v := state.Max()
	// fmt.Printf("getBest v( %v", v)
	// if v != nil {
	// 	val := v.(deque.Deque)
	// 	// fmt.Printf("getBest val( %s )\n", val)
	// 	if val == nil {
	// 		panic("getBest val == nil")
	// 	}
	// 	if val.Empty() == false {
	// 		b := val.Front()
	// 		if b == nil {
	// 			panic("getBest val.Front == nil")
	// 		}
	// 		best := b.(*Library)
	// 		return best
	// 	}
	// }
	// return nil
	return out
	// return val
}
