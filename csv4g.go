package csv4g

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type FieldDefine struct {
	reflect.StructField
	FieldIndex int
}

type Csv4g struct {
	name       string
	fields     []*FieldDefine
	lines      [][]string
	lineNo     int
	LineLen    int
	lineOffset int
}

type Option struct {
	Comma      rune
	LazyQuotes bool
	SkipLine   int
}

func Comma(r rune) func(*Option)      { return func(opt *Option) { opt.Comma = r } }
func LazyQuotes(b bool) func(*Option) { return func(opt *Option) { opt.LazyQuotes = b } }
func SkipLine(l int) func(*Option)    { return func(opt *Option) { opt.SkipLine = l } }

// Deprecated, Please use NewWithOpts
func New(filePath string, comma rune, lazyQuotes bool, o interface{}, skipLine int) (*Csv4g, error) {
	file, openErr := os.Open(filePath)
	if openErr != nil {
		return nil, fmt.Errorf("%s open file error %v", file.Name(), openErr)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = comma
	r.LazyQuotes = lazyQuotes
	var err error
	var fields []string
	fields, err = r.Read() // first line is field's description
	if err != nil {
		return nil, fmt.Errorf("%s read first line error %v", file.Name(), err)
	}
	offset := skipLine
	for skipLine > 0 {
		skipLine--
		_, err = r.Read()
		if err != nil {
			return nil, fmt.Errorf("%s skip line error %v", file.Name(), err)
		}
	}

	tType := reflect.TypeOf(o)
	if tType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%v must be a struct, cannot be an interface or pointer", tType.Elem().Name())
	}
	ret := &Csv4g{
		name:       file.Name(),
		fields:     make([]*FieldDefine, 0),
		lineNo:     0,
		lineOffset: offset + 1}

Out:
	for i := 0; i < tType.NumField(); i++ {
		f := tType.Field(i)
		tagStr := f.Tag.Get("csv")
		fieldName := f.Name
		canSkip := false
		if tagStr != "" {
			tags := strings.Split(tagStr, ",")
			for _, tag := range tags {
				switch tag {
				case "-":
					continue Out
				case "omitempty":
					canSkip = true
				default:
					fieldName = tag
				}
			}
		}
		fd := &FieldDefine{f, 0}
		index := -1
		for j, _ := range fields {
			if fields[j] == fieldName {
				index = j
				break
			}
		}
		if index == -1 {
			if !canSkip {
				return nil, fmt.Errorf("%s cannot find field %s", file.Name(), f.Name)
			}
			continue
		}
		fd.FieldIndex = index
		ret.fields = append(ret.fields, fd)
	}

	var lines [][]string
	lines, err = r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s Read error %v", file.Name(), err)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("%s has no data!", file.Name())
	}
	ret.lines = lines
	ret.LineLen = len(lines)
	return ret, nil
}

func NewWithOpts(filePath string, o interface{}, options ...func(*Option)) (*Csv4g, error) {
	file, openErr := os.Open(filePath)
	if openErr != nil {
		return nil, fmt.Errorf("%s open file error %v", file.Name(), openErr)
	}
	defer file.Close()

	defaultOpt := &Option{Comma: ',', LazyQuotes: false, SkipLine: 0}
	for _, opt := range options {
		opt(defaultOpt)
	}

	r := csv.NewReader(file)
	r.Comma = defaultOpt.Comma
	r.LazyQuotes = defaultOpt.LazyQuotes
	var err error
	var fields []string
	fields, err = r.Read() // first line is field's description
	if err != nil {
		return nil, fmt.Errorf("%s read first line error %v", file.Name(), err)
	}
	offset := defaultOpt.SkipLine
	for i := offset; i > 0; {
		i--
		_, err = r.Read()
		if err != nil {
			return nil, fmt.Errorf("%s skip line error %v", file.Name(), err)
		}
	}

	tType := reflect.TypeOf(o)
	if tType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%v must be a struct, cannot be an interface or pointer", tType.Elem().Name())
	}
	ret := &Csv4g{
		name:       file.Name(),
		fields:     make([]*FieldDefine, 0),
		lineNo:     0,
		lineOffset: offset + 1}

Out:
	for i := 0; i < tType.NumField(); i++ {
		f := tType.Field(i)
		tagStr := f.Tag.Get("csv")
		fieldName := f.Name
		canSkip := false
		if tagStr != "" {
			tags := strings.Split(tagStr, ",")
			for _, tag := range tags {
				switch tag {
				case "-":
					continue Out
				case "omitempty":
					canSkip = true
				default:
					fieldName = tag
				}
			}
		}
		fd := &FieldDefine{f, 0}
		index := -1
		for j, _ := range fields {
			if fields[j] == fieldName {
				index = j
				break
			}
		}
		if index == -1 {
			if !canSkip {
				return nil, fmt.Errorf("%s cannot find field %s", file.Name(), f.Name)
			}
			continue
		}
		fd.FieldIndex = index
		ret.fields = append(ret.fields, fd)
	}

	var lines [][]string
	lines, err = r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s read error %v", file.Name(), err)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("%s has no data!", file.Name())
	}
	ret.lines = lines
	ret.LineLen = len(lines)
	return ret, nil
}

func (this *Csv4g) Parse(obj interface{}) (err error) {
	if this.lineNo >= len(this.lines) {
		return io.EOF
	}
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%s error on parse line %d [%v]", this.name, this.lineNo+this.lineOffset+1, x)
			return
		}
		this.lineNo++
	}()
	values := this.lines[this.lineNo]
	elem := reflect.ValueOf(obj).Elem()
	for index, _ := range this.fields {
		f := elem.FieldByIndex(this.fields[index].Index)
		value := values[this.fields[index].FieldIndex]
		if conv, ok := converters[f.Kind()]; ok {
			v := conv(value)
			f.Set(v)
		} else {
			if f.Kind() == reflect.Slice {
				if sliceConv, ok := sliceConvertes[f.Type()]; ok {
					f.Set(sliceConv(strings.Split(value, "|")))
				} else {
					err = fmt.Errorf("%s:[%d] unsupported field set %v -> %v :[%d].",
						this.name, this.lineNo+this.lineOffset, this.fields[index], value)
				}
			} else {
				err = fmt.Errorf("%s:[%d] unsupported field set %v -> %v :[%d].",
					this.name, this.lineNo+this.lineOffset, this.fields[index], value)
			}
		}
	}
	return
}
