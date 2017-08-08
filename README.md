csv4g
=======

A csv file mapping to struct tool.

installation
------------

    go get github.com/cuixin/csv4g

example
-------

```
package main

import (
	"fmt"

	"github.com/cuixin/csv4g"
)

type Test struct {
	Id           int
	Name         string
	Desc         string
	Go           string
	Num          float32
	Foo          bool
	SliceInt     []int
	SliceFloat32 []float32 `csv:"sliceFloat32"`
	IgnoreField  string    `csv:"-"`
	CustomField  string    `csv:"custom,omitempty"`
	EmptyField   string    `csv:"omitempty"`
}

func main() {
	csv, err := csv4g.NewWithOpts("test.csv", Test{}, csv4g.Comma(','), csv4g.LazyQuotes(true), csv4g.SkipLine(1))
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
	for i := 0; i < csv.LineLen; i++ {
		tt := &Test{}
		err = csv.Parse(tt)
		if err != nil {
			fmt.Printf("Error on parse %v\n", err)
			return
		}
		fmt.Println(tt)
	}
}


```
