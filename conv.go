package csv4g

import (
	"reflect"
	"strconv"
)

var invalidValue = reflect.Value{}

// Default converters for basic types.
var converters = map[reflect.Kind]func(string) reflect.Value{
	reflect.Bool:    convertBool,
	reflect.Float32: convertFloat32,
	reflect.Float64: convertFloat64,
	reflect.Int:     convertInt,
	reflect.Int8:    convertInt8,
	reflect.Int16:   convertInt16,
	reflect.Int32:   convertInt32,
	reflect.Int64:   convertInt64,
	reflect.String:  convertString,
	reflect.Uint:    convertUint,
	reflect.Uint8:   convertUint8,
	reflect.Uint16:  convertUint16,
	reflect.Uint32:  convertUint32,
	reflect.Uint64:  convertUint64,
}
var (
	boolSliceType    = reflect.TypeOf([]bool{})
	float32SliceType = reflect.TypeOf([]float32{})
	float64SliceType = reflect.TypeOf([]float64{})
	intSliceType     = reflect.TypeOf([]int{})
	int8SliceType    = reflect.TypeOf([]int8{})
	int16SliceType   = reflect.TypeOf([]int16{})
	int32SliceType   = reflect.TypeOf([]int32{})
	int64SliceType   = reflect.TypeOf([]int64{})
	stringSliceType  = reflect.TypeOf([]string{})
	uintSliceType    = reflect.TypeOf([]uint{})
	uint8SliceType   = reflect.TypeOf([]uint8{})
	uint16SliceType  = reflect.TypeOf([]uint16{})
	uint32SliceType  = reflect.TypeOf([]uint32{})
	uint64SliceType  = reflect.TypeOf([]uint64{})
)

var sliceConvertes = map[reflect.Type]func([]string) reflect.Value{
	boolSliceType:    convertBools,
	float32SliceType: convertFloat32s,
	float64SliceType: convertFloat64s,
	intSliceType:     convertInts,
	int8SliceType:    convertInt8s,
	int16SliceType:   convertInt16s,
	int32SliceType:   convertInt32s,
	int64SliceType:   convertInt64s,
	stringSliceType:  convertStrings,
	uint8SliceType:   convertUint8s,
	uint16SliceType:  convertUint16s,
	uint32SliceType:  convertUint32s,
	uint64SliceType:  convertUint64s,
}

func convertBool(value string) reflect.Value {
	if v, err := strconv.ParseBool(value); err == nil {
		return reflect.ValueOf(v)
	}
	return invalidValue
}

func convertBools(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertBool(values[i])
	}
	ret := reflect.MakeSlice(boolSliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertFloat32(value string) reflect.Value {
	if v, err := strconv.ParseFloat(value, 32); err == nil {
		return reflect.ValueOf(float32(v))
	}
	return invalidValue
}

func convertFloat32s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertFloat32(values[i])
	}
	ret := reflect.MakeSlice(float32SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertFloat64(value string) reflect.Value {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		return reflect.ValueOf(v)
	}
	return invalidValue
}

func convertFloat64s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertFloat64(values[i])
	}
	ret := reflect.MakeSlice(float64SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertInt(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 0); err == nil {
		return reflect.ValueOf(int(v))
	}
	return invalidValue
}

func convertInts(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertInt(values[i])
	}
	ret := reflect.MakeSlice(intSliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertInt8(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 8); err == nil {
		return reflect.ValueOf(int8(v))
	}
	return invalidValue
}

func convertInt8s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertInt8(values[i])
	}
	ret := reflect.MakeSlice(int8SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertInt16(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 16); err == nil {
		return reflect.ValueOf(int16(v))
	}
	return invalidValue
}

func convertInt16s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertInt16(values[i])
	}
	ret := reflect.MakeSlice(int16SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertInt32(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 32); err == nil {
		return reflect.ValueOf(int32(v))
	}
	return invalidValue
}

func convertInt32s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertInt32(values[i])
	}
	ret := reflect.MakeSlice(int32SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertInt64(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 64); err == nil {
		return reflect.ValueOf(v)
	}
	return invalidValue
}

func convertInt64s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertInt64(values[i])
	}
	ret := reflect.MakeSlice(int64SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertString(value string) reflect.Value {
	return reflect.ValueOf(value)
}

func convertStrings(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertString(values[i])
	}
	ret := reflect.MakeSlice(stringSliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertUint(value string) reflect.Value {
	if v, err := strconv.ParseUint(value, 10, 0); err == nil {
		return reflect.ValueOf(uint(v))
	}
	return invalidValue
}

func convertUInts(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertUint(values[i])
	}
	ret := reflect.MakeSlice(uintSliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertUint8(value string) reflect.Value {
	if v, err := strconv.ParseUint(value, 10, 8); err == nil {
		return reflect.ValueOf(uint8(v))
	}
	return invalidValue
}
func convertUint8s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertUint8(values[i])
	}
	ret := reflect.MakeSlice(uint8SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertUint16(value string) reflect.Value {
	if v, err := strconv.ParseUint(value, 10, 16); err == nil {
		return reflect.ValueOf(uint16(v))
	}
	return invalidValue
}

func convertUint16s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertUint16(values[i])
	}
	ret := reflect.MakeSlice(uint16SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertUint32(value string) reflect.Value {
	if v, err := strconv.ParseUint(value, 10, 32); err == nil {
		return reflect.ValueOf(uint32(v))
	}
	return invalidValue
}

func convertUint32s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertUint32(values[i])
	}
	ret := reflect.MakeSlice(uint32SliceType, 0, 0)
	return reflect.Append(ret, items...)
}

func convertUint64(value string) reflect.Value {
	if v, err := strconv.ParseUint(value, 10, 64); err == nil {
		return reflect.ValueOf(v)
	}
	return invalidValue
}

func convertUint64s(values []string) reflect.Value {
	items := make([]reflect.Value, len(values))
	for i, _ := range values {
		items[i] = convertUint64(values[i])
	}
	ret := reflect.MakeSlice(uint64SliceType, 0, 0)
	return reflect.Append(ret, items...)
}
