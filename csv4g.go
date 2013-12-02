package csv4g

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "os"
    "reflect"
    "strconv"
)

type Csv4g struct {
    name     string
    fieldMap map[*reflect.Value]int
    lines    [][]string
    lineNo   int
    LineLen  int
}

const lineOffset = 2

func New(filePath string, comma rune, o interface{}) (*Csv4g, error) {
    file, openErr := os.Open(filePath)
    if openErr != nil {
        return nil, openErr
    }
    defer file.Close()
    r := csv.NewReader(file)
    r.Comma = comma
    fields, err := r.Read()
    if err != nil {
        return nil, err
    }
    value := reflect.ValueOf(o)
    fieldMap, err := checkFields(fields, &value, file.Name())
    if err != nil {
        return nil, errors.New(fmt.Sprintf("%s %s", file.Name(), err.Error()))
    }
    var lines [][]string
    lines, err = r.ReadAll()
    if err != nil {
        return nil, err
    }
    if len(lines) == 0 {
        return nil, errors.New(fmt.Sprintf("%s has no data!", file.Name()))
    }
    return &Csv4g{name: file.Name(),
        fieldMap: fieldMap,
        lines:    lines, lineNo: 0, LineLen: len(lines)}, nil
}

func checkFields(fields []string, v *reflect.Value, name string) (map[*reflect.Value]int, error) {
    fm := make(map[*reflect.Value]int)
    e := v.Elem()
    for k, v := range fields {
        f := e.FieldByName(v)
        if !f.IsValid() {
            return nil, errors.New(fmt.Sprintf("%s cannot find field %s", name, f))
        }
        fm[&f] = k
    }

    csv_size := len(fm)
    struct_size := e.NumField()
    if csv_size < struct_size {
        return nil, errors.New(fmt.Sprintf(
            "%s field size is not equal, csv = %d, struct = %d", name,
            csv_size, struct_size))
    }
    return fm, nil
}

func (this *Csv4g) Parse(obj interface{}) error {
    if this.lineNo >= len(this.lines) {
        return io.EOF
    }
    defer func() { this.lineNo++ }()
    values := this.lines[this.lineNo]
    elem := reflect.ValueOf(obj).Elem()
    for field, index := range this.fieldMap {
        f := elem.Field(index)
        switch f.Kind() {
        case reflect.Bool:
            b, err := strconv.ParseBool(values[index])
            if err != nil {
                return fmt.Errorf("%s:[%d] %v", this.name, this.lineNo+lineOffset, err)
            }
            f.SetBool(b)
        case reflect.Float32, reflect.Float64:
            f64, err := strconv.ParseFloat(values[index], 64)
            if err != nil {
                return fmt.Errorf("%s:[%d] %v", this.name, this.lineNo+lineOffset, err)
            }
            f.SetFloat(f64)
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            i64, err := strconv.ParseInt(values[index], 10, 64)
            if err != nil {
                return fmt.Errorf("%s:[%d] %v", this.name, this.lineNo+lineOffset, err)
            }
            f.SetInt(i64)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            ui64, err := strconv.ParseUint(values[index], 10, 64)
            if err != nil {
                return fmt.Errorf("%s:[%d] %v", this.name, this.lineNo+lineOffset, err)
            }
            f.SetUint(ui64)
        case reflect.String:
            f.SetString(values[index])
        default:
            return fmt.Errorf("%s:[%d] unsupported field set %s -> %v :[%d].",
                this.name, this.lineNo+lineOffset, field, values[index])
        }
    }

    return nil
}
