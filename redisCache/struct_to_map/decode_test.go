package struct_to_map

import (
	"reflect"
	"testing"
)

func Test_decodeStruct(t *testing.T) {
	var testStruct = Request{}
	var f = reflect.ValueOf(&testStruct.Ping)
	f = f.Elem()
	err := decodeString(f, "张三")
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!! ts:%+v", testStruct.Ping)
	}
}
