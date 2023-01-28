package struct_to_map

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func checkType(value reflect.Value) bool {
	switch value.Interface().(type) { // redis只支持存储基本类型数据，复杂类型数据需要序列化成[]byte类型
	case bool,
		int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64,
		float32,
		float64,
		[]byte,
		string:
		return true
	default:
		return false
	}
}

func Encode(data interface{}) []interface{} {
	var typeOf = reflect.TypeOf(data)
	var valueOf = reflect.ValueOf(data)
	if typeOf.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	}
	var length = typeOf.NumField()
	var filter = make([]interface{}, length*2)
	for i := 0; i < length; i++ {
		filter[i*2] = typeOf.Field(i).Name
		var value = valueOf.FieldByName(typeOf.Field(i).Name)
		if !checkType(value) {
			js, err := json.Marshal(value.Interface())
			if err != nil {
				fmt.Println(err)
			}
			filter[i*2+1] = js
		} else {
			filter[i*2+1] = value.Interface()
		}
	}
	return filter
}
