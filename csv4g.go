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

const (
    ToMap = iota
    ToArray
)

func checkFields(fieldMap map[string]int, elem *reflect.Value) error {
    for f, _ := range fieldMap {
        if !elem.FieldByName(f).IsValid() {
            return errors.New(fmt.Sprintf("cannot find field %s", f))
        }
    }
    return nil
}

func setValue(toData interface{}, fields map[string]int,
    values []string, elem *reflect.Value) error {
    for field, index := range fields {
        f := elem.FieldByName(field)
        switch f.Kind() {
        case reflect.Bool:
            b, err := strconv.ParseBool(values[index])
            if err != nil {
                return err
            }
            f.SetBool(b)
        case reflect.Float32:
            f32, err := strconv.ParseFloat(values[index], 32)
            if err != nil {
                return err
            }
            f.SetFloat(f32)
        case reflect.Float64:
            f64, err := strconv.ParseFloat(values[index], 64)
            if err != nil {
                return err
            }
            f.SetFloat(f64)
        case reflect.Int8:
            i, err := strconv.ParseInt(values[index], 10, 8)
            if err != nil {
                return err
            }
            f.SetInt(i)
        case reflect.Int16:
            i, err := strconv.ParseInt(values[index], 10, 16)
            if err != nil {
                return err
            }
            f.SetInt(i)
        case reflect.Int32:
            i, err := strconv.ParseInt(values[index], 10, 32)
            if err != nil {
                return err
            }
            f.SetInt(i)
        case reflect.Int64:
            i, err := strconv.ParseInt(values[index], 10, 64)
            if err != nil {
                return err
            }
            f.SetInt(i)
        case reflect.Uint8:
            i, err := strconv.ParseUint(values[index], 10, 8)
            if err != nil {
                return err
            }
            f.SetUint(i)
        case reflect.Uint16:
            i, err := strconv.ParseUint(values[index], 10, 16)
            if err != nil {
                return err
            }
            f.SetUint(i)
        case reflect.Uint32:
            i, err := strconv.ParseUint(values[index], 10, 32)
            if err != nil {
                return err
            }
            f.SetUint(i)
        case reflect.Uint64:
            i, err := strconv.ParseUint(values[index], 10, 64)
            if err != nil {
                return err
            }
            f.SetUint(i)

        case reflect.String:
            f.SetString(values[index])
        }
    }
    if w, ok := toData.(map[string]interface{}); ok {
        w[values[0]] = elem.Interface()
    }
    return nil
}

func Parse(filePath string, nilType interface{}, toType int) (interface{}, error) {
    file, openErr := os.Open(filePath)
    if openErr != nil {
        return nil, openErr
    }
    defer file.Close()
    r := csv.NewReader(file)
    fields, err := r.Read()
    if err != nil {
        return nil, err
    }
    fieldMap := make(map[string]int)

    for k, v := range fields {
        fieldMap[v] = k
    }
    nilObj := reflect.TypeOf(nilType)
    objPtr := reflect.New(nilObj)
    elem := objPtr.Elem()

    csv_size := len(fieldMap)
    struct_size := elem.NumField()

    if csv_size != struct_size {
        return nil, errors.New(fmt.Sprintf(
            "%s's field size is not equal, csv = %d, struct = %d",
            file.Name(), csv_size, struct_size))
    }

    err = checkFields(fieldMap, &elem)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("%s %s", file.Name(), err.Error()))
    }
    toData := make(map[string]interface{})

    for {
        values, err := r.Read()
        if err != nil {
            if err != io.EOF {
                return nil, err
            }
            if toType == ToMap {
                arrData := make([]interface{}, len(toData))
                for k, v := range toData {
                    i, _ := strconv.ParseInt(k, 10, 32)
                    arrData[i] = v
                }
                return arrData, nil
            } else {
                return toData, nil
            }
        }
        err = setValue(toData, fieldMap, values, &elem)
        if err != nil {
            return nil, errors.New(fmt.Sprintf("%s parse data %s", file.Name(), err.Error()))
        }
    }
    return nil, errors.New(fmt.Sprintf("%s has no data!", file.Name()))
}
