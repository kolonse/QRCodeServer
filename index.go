package main

import (
	"strconv"
	"sync"
)

var m *sync.Mutex
var str64 = "jyGu9rOINxWoMHQ3zn8mLAgF56qaMviYSfPEXJelwb7hT4pUt0sKRdcBVCDZ21k"

func makeUnique() string {
	return toStr64(readAutoInc())
}

func readAutoInc() int {
	m.Lock()
	defer m.Unlock()
	v := kvAutoIncrease.Read("autoincrease").ToInt()
	v += 1
	kvAutoIncrease.Write("autoincrease", []byte(strconv.Itoa(v)))
	return v
}

func toStr64(v int) string {
	ret := make([]byte, 0)
	for {
		mod := v % len(str64)
		ret = append(ret, str64[mod])
		v = (v - mod) / len(str64)
		if v == 0 {
			break
		}
	}
	return string(ret)
}

func init() {
	m = new(sync.Mutex)
}
