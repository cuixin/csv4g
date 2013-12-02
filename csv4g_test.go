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
    // }
    csv, err := New("test.csv", ',', &Test{})
    if err != nil {
        t.Errorf("Error %v\n", err)
        return
    }
    for i := 0; i < csv.LineLen; i++ {
        tt := &Test{}
        err = csv.Parse(tt)
        if err != nil {
            t.Errorf("Error on parse %v\n", err)
        }
        fmt.Println(tt)
    }
}
