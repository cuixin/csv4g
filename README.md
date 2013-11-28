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
        Id   int
        Name string
        Desc string
        Go   string
        Num  float32
        Foo  bool
    }

    func main() {
        data, err := csv4g.Parse("./csv4g/test.csv", Test{}, csv4g.ToMap)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(data)
    }

```
