package struct_to_map

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// decoderFunc represents decoding functions for default built-in types.
type decoderFunc func(reflect.Value, string) error

var (
	// List of built-in decoders indexed by their numeric constant values (eg: reflect.Bool = 1).
	decoders = []decoderFunc{
		reflect.Bool:          decodeBool,
		reflect.Int:           decodeInt,
		reflect.Int8:          decodeInt8,
		reflect.Int16:         decodeInt16,
		reflect.Int32:         decodeInt32,
		reflect.Int64:         decodeInt64,
		reflect.Uint:          decodeUint,
		reflect.Uint8:         decodeUint8,
		reflect.Uint16:        decodeUint16,
		reflect.Uint32:        decodeUint32,
		reflect.Uint64:        decodeUint64,
		reflect.Float32:       decodeFloat32,
		reflect.Float64:       decodeFloat64,
		reflect.Complex64:     decodeJson,
		reflect.Complex128:    decodeJson,
		reflect.Array:         decodeJson,
		reflect.Chan:          decodeJson,
		reflect.Func:          decodeJson,
		reflect.Interface:     decodeJson,
		reflect.Map:           decodeJson,
		reflect.Ptr:           decodeJson,
		reflect.Slice:         decodeSlice,
		reflect.String:        decodeString,
		reflect.Struct:        decodeJson,
		reflect.UnsafePointer: decodeJson,
	}

	// Global map of struct field specs that is populated once for every new
	// struct type that is scanned. This caches the field types and the corresponding
	// decoder functions to avoid iterating through struct fields on subsequent scans.
	globalStructMap = newStructMap()
)

func Struct(dst interface{}) (StructValue, error) {
	v := reflect.ValueOf(dst)

	// The destination to scan into should be a struct pointer.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return StructValue{}, fmt.Errorf("redis.Scan(non-pointer %T)", dst)
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return StructValue{}, fmt.Errorf("redis.Scan(non-struct %T)", dst)
	}

	return StructValue{
		spec:  globalStructMap.get(v.Type()),
		value: v,
	}, nil
}

func Decode(dst interface{}, data map[string]string) error {

	strct, err := Struct(dst)
	if err != nil {
		return err
	}

	// Iterate through the (key, value) sequence.
	for key, val := range data {
		if err := strct.Scan(key, val); err != nil {
			return err
		}
	}
	return nil
}

func decodeBool(f reflect.Value, s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	f.SetBool(b)
	return nil
}

func decodeInt8(f reflect.Value, s string) error {
	return decodeNumber(f, s, 8)
}

func decodeInt16(f reflect.Value, s string) error {
	return decodeNumber(f, s, 16)
}

func decodeInt32(f reflect.Value, s string) error {
	return decodeNumber(f, s, 32)
}

func decodeInt64(f reflect.Value, s string) error {
	return decodeNumber(f, s, 64)
}

func decodeInt(f reflect.Value, s string) error {
	return decodeNumber(f, s, 0)
}

func decodeNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetInt(v)
	return nil
}

func decodeUint8(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 8)
}

func decodeUint16(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 16)
}

func decodeUint32(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 32)
}

func decodeUint64(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 64)
}

func decodeUint(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 0)
}

func decodeUnsignedNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetUint(v)
	return nil
}

func decodeFloat32(f reflect.Value, s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	f.SetFloat(v)
	return nil
}

// although the default is float64, but we better define it.
func decodeFloat64(f reflect.Value, s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.SetFloat(v)
	return nil
}

func decodeString(f reflect.Value, s string) error {
	f.SetString(s)
	return nil
}

func decodeSlice(f reflect.Value, s string) error {
	// []byte slice ([]uint8).
	if f.Type().Elem().Kind() == reflect.Uint8 {
		f.SetBytes([]byte(s))
	} else {
		decodeJson(f, s)
	}
	return nil
}

func decodeJson(f reflect.Value, s string) error {
	var i = f.Interface()
	rType := reflect.TypeOf(i)
	newElem := reflect.New(rType)
	i = newElem.Interface()
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		return err
	}
	newElem = newElem.Elem()
	f.Set(newElem)
	return nil
}

func decodeUnsupported(v reflect.Value, s string) error {
	return fmt.Errorf("redis.Scan(unsupported %s)", v.Type())
}