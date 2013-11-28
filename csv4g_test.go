package csv4g

import "testing"
import "fmt"

type Test struct {
    Id   int
    Name string
    Desc string
    Go   string
    Num  float32
    Foo  bool
}

func TestParse(t *testing.T) {
    data, err := Parse("test.csv", Test{}, ToArray)
    fmt.Println(data)
    if err != nil {
        t.Errorf("Error %v\n", err)
        return
    }
    t.Logf("Result %v\n", data)
}
