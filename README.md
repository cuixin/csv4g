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

import "github.com/cuixin/csv4g"
import "fmt"

type Test struct {
	Id           int
	Name         string
	Desc         string
	Go           string
	Num          float32
	Foo          bool
	SliceInt     []string
	SliceFloat32 []string
	IgnoreField  string `csv:"-"`
}

func main() {
	comma := func(opt *csv4g.Option) {
		opt.Comma = ',' // default is comma
	}
	lazyQuotes := func(opt *csv4g.Option) {
		opt.LazyQuotes = true // default is false
	}
	skipLine := func(opt *csv4g.Option) {
		opt.SkipLine = 1 // default is 0
	}
	csv, err := csv4g.NewWithOpts("test.csv", Test{}, comma, lazyQuotes, skipLine)
	if err != nil {
		fmt.Errorf("Error %v\n", err)
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
