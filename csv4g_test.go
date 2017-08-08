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
	IgnoreField  string `csv:"-"`
	CustomField  string `csv:"custom,omitempty"`
	EmptyField   string `csv:"omitempty"`
}

func TestParse(t *testing.T) {
	testFiles := []string{"test.csv", "test_empty.csv"}
	for _, testFile := range testFiles {
		csv, err := New(testFile, ',', true, Test{}, 1)
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
}

func TestParseWithOptions(t *testing.T) {
	testFiles := []string{"test.csv", "test_empty.csv"}
	for _, testFile := range testFiles {
		comma := func(opt *Option) {
			opt.Comma = ','
		}
		lazyQuotes := func(opt *Option) {
			opt.LazyQuotes = true
		}
		skipLine := func(opt *Option) {
			opt.SkipLine = 1
		}
		csv, err := NewWithOpts(testFile, Test{}, comma, lazyQuotes, skipLine)
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
}
