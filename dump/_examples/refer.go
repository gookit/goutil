package main

import (
	"github.com/kortschak/utter"
	"github.com/kr/pretty"
)

func main() {
	vs := []interface{}{123}

	// print var data
	_, err := pretty.Println(vs)
	if err != nil {
		panic(err)
	}

	// print var data
	for _, v := range vs {
		utter.Dump(v)
	}
}
