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
    if testArrayData, ok := data.([]Test); ok {
        fmt.Println(testArrayData[1])
    }

    data2, err2 := Parse("test.csv", Test{}, ToMap)
    fmt.Println(data2)
    if err2 != nil {
        t.Errorf("Error %v\n", err2)
        return
    }

    if testMapData, ok := data2.(map[string]interface{}); ok {
        fmt.Println(testMapData["1"])
    }
}
