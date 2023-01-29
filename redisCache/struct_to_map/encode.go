package struct_to_map

import (
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"reflect"
	"unicode"
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
	var filter = make([]interface{}, 0)
	var ok bool
	var fieldName string
	for i := 0; i < length; i++ {
		fieldName = typeOf.Field(i).Name
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}
		var value = valueOf.Field(i)
		if !checkType(value) {
			var itf = value.Interface()
			var js []byte
			var err error
			var pm proto.Message
			if pm, ok = itf.(proto.Message); ok {
				if itf != nil && !value.IsNil() {
					js, err = proto.Marshal(pm)
				} else {
					js = Nil
				}
			} else {
				js, err = json.Marshal(itf)
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			filter = append(filter, GetTagByField(typeOf.Field(i), TagName))
			filter = append(filter, js)
		} else {
			filter = append(filter, GetTagByField(typeOf.Field(i), TagName))
			filter = append(filter, value.Interface())
		}
		//filter[i*2] = typeOf.Field(i).Name
	}
	return filter
}
