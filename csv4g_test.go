package csv4g

import "testing"
import "fmt"

type Test struct {
	Id           int
	Name         string
	Desc         string
	Go           string
	Num          float32
	Foo          bool
	SliceInt     []int
	SliceFloat32 []float32
}

func TestParse(t *testing.T) {
	csv, err := New("test.csv", ',', Test{}, 1)
	if err != nil {
		t.Errorf("Error %v\n", err)
		return
	}
	// fmt.Println(csv)
	for i := 0; i < csv.LineLen; i++ {
		tt := &Test{}
		err = csv.Parse(tt)
		if err != nil {
			t.Errorf("%v\n", err)
			break
		}
		fmt.Println(tt)
	}

}
