package main

import (
	//	"math/rand"
	//	"sort"
	"testing"
)

//type StrSort []byte

//func (a StrSort) Len() int      { return len(a) }
//func (a StrSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
//func (a StrSort) Less(i, j int) bool {
//	if rand.Intn(2) == 0 {
//		return false
//	}
//	return true
//}

//func (s str64)
func TestUnique(t *testing.T) {
	//	s := make([]byte, len(str64))
	//	copy(s, []byte(str64))
	//	sort.Sort(StrSort(s))
	//	t.Log(string(s))
	t.Log(makeUnique())
	t.Log(makeUnique())
	t.Log(makeUnique())
	t.Log(makeUnique())
	t.Log(makeUnique())
	t.Log(makeUnique())
}
