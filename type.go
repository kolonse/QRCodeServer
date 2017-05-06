package main

import (
	"github.com/kolonse/simplekv"
)

type Shortest struct {
	Redirect bool
	Content  string
	Index    string
}

var kv *simplekv.SKV
var kvIndex2MD5 *simplekv.SKV
var kvAutoIncrease *simplekv.SKV
